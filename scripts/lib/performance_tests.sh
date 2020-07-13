function check_for_timeout() {
    if [[ $((SECONDS - $1)) -gt 1500 ]]; then
        FAILURE_COUNT=$((FAILURE_COUNT + 1))
        HAS_FAILED=true
        if [[ $FAILURE_COUNT -gt 1 ]]; then
            RUNNING_PERFORMANCE=false
            echo "Failed twice, deprovisioning cluster"
            on_error
        fi
        echo "Failed once, retrying"
    fi
}

function save_results_to_file() {
    echo $filename
    echo "Date", "\"slot1\"", "\"slot2\"" >> $filename
    echo $(date +"%m-%d-%Y-%T"), $((SCALE_UP_DURATION_ARRAY[0])), $((SCALE_DOWN_DURATION_ARRAY[0])) >> $filename
    echo $(date +"%m-%d-%Y-%T"), $((SCALE_UP_DURATION_ARRAY[1])), $((SCALE_DOWN_DURATION_ARRAY[1])) >> $filename
    echo $(date +"%m-%d-%Y-%T"), $((SCALE_UP_DURATION_ARRAY[2])), $((SCALE_DOWN_DURATION_ARRAY[2])) >> $filename

    cat $filename
    if [[ ${#PERFORMANCE_TEST_S3_BUCKET_NAME} -gt 0 ]]; then
        aws s3 cp $filename ${PERFORMANCE_TEST_S3_BUCKET_NAME}${1}
    else
        echo "No S3 bucket name given, skipping test result upload."
    fi
}

function run_performance_test_130_pods() {
    echo "Running performance tests against cluster"
    RUNNING_PERFORMANCE=true
    $KUBECTL_PATH apply -f ./testdata/deploy-130-pods.yaml

    DEPLOY_START=$SECONDS
    FAILURE_COUNT=0

    SCALE_UP_DURATION_ARRAY=()
    SCALE_DOWN_DURATION_ARRAY=()
    while [ ${#SCALE_DOWN_DURATION_ARRAY[@]} -lt 3 ]
    do
        ITERATION_START=$SECONDS
        HAS_FAILED=false
        $KUBECTL_PATH scale -f ./testdata/deploy-130-pods.yaml --replicas=130
        while [[ ! $($KUBECTL_PATH get deploy | grep 130/130) && "$HAS_FAILED" == false ]]
        do
            sleep 1
            echo "Scaling UP"
            echo $($KUBECTL_PATH get deploy)
            check_for_timeout $ITERATION_START
        done

        if [[ "$HAS_FAILED" == false ]]; then
            SCALE_UP_DURATION_ARRAY+=( $((SECONDS - ITERATION_START)) )
        fi
        MIDPOINT_START=$SECONDS
        $KUBECTL_PATH scale -f ./testdata/deploy-130-pods.yaml --replicas=0
        while [[ $($KUBECTL_PATH get pods) ]]
        do
            sleep 1
            echo "Scaling DOWN"
            echo $($KUBECTL_PATH get deploy)
        done
        if [[ "$HAS_FAILED" == false ]]; then
            SCALE_DOWN_DURATION_ARRAY+=($((SECONDS - MIDPOINT_START)))
        fi
    done

    echo "Times to scale up:"
    INDEX=0
    while [ $INDEX -lt ${#SCALE_UP_DURATION_ARRAY[@]} ]
    do
        echo ${SCALE_UP_DURATION_ARRAY[$INDEX]}
        INDEX=$((INDEX + 1))
    done
    echo ""
    echo "Times to scale down:"
    INDEX=0
    while [ $INDEX -lt ${#SCALE_DOWN_DURATION_ARRAY[@]} ]
    do
        echo "${SCALE_DOWN_DURATION_ARRAY[$INDEX]} seconds"
        INDEX=$((INDEX + 1))
    done
    echo ""
    DEPLOY_DURATION=$((SECONDS - DEPLOY_START))

    filename="pod-130-Test#${TEST_ID}-$(date +"%m-%d-%Y-%T")-${TEST_IMAGE_VERSION}.csv"
    save_results_to_file "/130-pods/"
    
    echo "TIMELINE: 130 Pod performance test took $DEPLOY_DURATION seconds."
    RUNNING_PERFORMANCE=false
    $KUBECTL_PATH delete -f ./testdata/deploy-130-pods.yaml
}

function run_performance_test_730_pods() {
    echo "Running performance tests against cluster"
    RUNNING_PERFORMANCE=true
    $KUBECTL_PATH apply -f ./testdata/deploy-730-pods.yaml

    DEPLOY_START=$SECONDS
    FAILURE_COUNT=0

    SCALE_UP_DURATION_ARRAY=()
    SCALE_DOWN_DURATION_ARRAY=()
    while [ ${#SCALE_DOWN_DURATION_ARRAY[@]} -lt 3 ]
    do
        ITERATION_START=$SECONDS
        HAS_FAILED=false
        $KUBECTL_PATH scale -f ./testdata/deploy-730-pods.yaml --replicas=730
        while [[ ! $($KUBECTL_PATH get deploy | grep 730/730) && "$HAS_FAILED" == false ]]
        do
            sleep 2
            echo "Scaling UP"
            echo $($KUBECTL_PATH get deploy)
            check_for_timeout $ITERATION_START
        done

        if [[ "$HAS_FAILED" == false ]]; then
            SCALE_UP_DURATION_ARRAY+=( $((SECONDS - ITERATION_START)) )
        fi
        MIDPOINT_START=$SECONDS
        $KUBECTL_PATH scale -f ./testdata/deploy-730-pods.yaml --replicas=0
        while [[ $($KUBECTL_PATH get pods) ]]
        do
            sleep 2
            echo "Scaling DOWN"
            echo $($KUBECTL_PATH get deploy)
        done
        if [[ "$HAS_FAILED" == false ]]; then
            SCALE_DOWN_DURATION_ARRAY+=($((SECONDS - MIDPOINT_START)))
        fi
    done

    echo "Times to scale up:"
    INDEX=0
    while [ $INDEX -lt ${#SCALE_UP_DURATION_ARRAY[@]} ]
    do
        echo ${SCALE_UP_DURATION_ARRAY[$INDEX]}
        INDEX=$((INDEX + 1))
    done
    echo ""
    echo "Times to scale down:"
    INDEX=0
    while [ $INDEX -lt ${#SCALE_DOWN_DURATION_ARRAY[@]} ]
    do
        echo "${SCALE_DOWN_DURATION_ARRAY[$INDEX]} seconds"
        INDEX=$((INDEX + 1))
    done
    echo ""
    DEPLOY_DURATION=$((SECONDS - DEPLOY_START))

    filename="pod-730-Test#${TEST_ID}-$(date +"%m-%d-%Y-%T")-${TEST_IMAGE_VERSION}.csv"
    save_results_to_file "/730-pods/"
    
    echo "TIMELINE: 730 Pod performance test took $DEPLOY_DURATION seconds."
    RUNNING_PERFORMANCE=false
    $KUBECTL_PATH delete -f ./testdata/deploy-730-pods.yaml
}

function scale_nodes_for_5000_pod_test() {
    AUTO_SCALE_GROUP_INFO=$(aws autoscaling describe-auto-scaling-groups | grep -B18 100,)
    echo "Group info ${AUTO_SCALE_GROUP_INFO}"
    AUTO_SCALE_GROUP_NAME_WITH_QUOTES=$(echo "${AUTO_SCALE_GROUP_INFO%%,*}")
    echo "Group name with quotes ${AUTO_SCALE_GROUP_NAME_WITH_QUOTES}"
    AUTO_SCALE_GROUP_NAME_WITH_QUOTES=$(echo "${AUTO_SCALE_GROUP_NAME_WITH_QUOTES##* }")
    echo "Group name with quotes ${AUTO_SCALE_GROUP_NAME_WITH_QUOTES}"
    AUTO_SCALE_GROUP_NAME="${AUTO_SCALE_GROUP_NAME_WITH_QUOTES%\"}"
    echo "Group name ${AUTO_SCALE_GROUP_NAME}"
    AUTO_SCALE_GROUP_NAME=$(echo $AUTO_SCALE_GROUP_NAME | cut -c2-)
    echo $AUTO_SCALE_GROUP_NAME

    aws autoscaling update-auto-scaling-group \
        --auto-scaling-group-name $AUTO_SCALE_GROUP_NAME \
        --desired-capacity 98
}

function run_performance_test_5000_pods() {
    echo "Running performance tests against cluster"
    RUNNING_PERFORMANCE=true
    $KUBECTL_PATH apply -f ./testdata/deploy-5000-pods.yaml
    
    DEPLOY_START=$SECONDS
    FAILURE_COUNT=0

    SCALE_UP_DURATION_ARRAY=()
    SCALE_DOWN_DURATION_ARRAY=()
    while [ ${#SCALE_DOWN_DURATION_ARRAY[@]} -lt 3 ]
    do
        ITERATION_START=$SECONDS
        HAS_FAILED=false
        $KUBECTL_PATH scale -f ./testdata/deploy-5000-pods.yaml --replicas=5000
        while [[ ! $($KUBECTL_PATH get deploy | grep 5000/5000) && "$HAS_FAILED" == false ]]
        do
            sleep 2
            echo "Scaling UP"
            echo $($KUBECTL_PATH get deploy)
            check_for_timeout $ITERATION_START
        done

        if [[ "$HAS_FAILED" == false ]]; then
            SCALE_UP_DURATION_ARRAY+=( $((SECONDS - ITERATION_START)) )
        fi
        MIDPOINT_START=$SECONDS
        $KUBECTL_PATH scale -f ./testdata/deploy-5000-pods.yaml --replicas=0
        while [[ $($KUBECTL_PATH get pods) ]]
        do
            sleep 2
            echo "Scaling DOWN"
            echo $($KUBECTL_PATH get deploy)
        done
        if [[ "$HAS_FAILED" == false ]]; then
            SCALE_DOWN_DURATION_ARRAY+=($((SECONDS - MIDPOINT_START)))
        fi
    done

    echo "Times to scale up:"
    INDEX=0
    while [ $INDEX -lt ${#SCALE_UP_DURATION_ARRAY[@]} ]
    do
        echo ${SCALE_UP_DURATION_ARRAY[$INDEX]}
        INDEX=$((INDEX + 1))
    done
    echo ""
    echo "Times to scale down:"
    INDEX=0
    while [ $INDEX -lt ${#SCALE_DOWN_DURATION_ARRAY[@]} ]
    do
        echo "${SCALE_DOWN_DURATION_ARRAY[$INDEX]} seconds"
        INDEX=$((INDEX + 1))
    done
    echo ""
    DEPLOY_DURATION=$((SECONDS - DEPLOY_START))

    filename="pod-5000-Test#${TEST_ID}-$(date +"%m-%d-%Y-%T")-${TEST_IMAGE_VERSION}.csv"
    save_results_to_file "/5000-pods/"
    
    echo "TIMELINE: 5000 Pod performance test took $DEPLOY_DURATION seconds."
    RUNNING_PERFORMANCE=false
    $KUBECTL_PATH delete -f ./testdata/deploy-5000-pods.yaml
}
