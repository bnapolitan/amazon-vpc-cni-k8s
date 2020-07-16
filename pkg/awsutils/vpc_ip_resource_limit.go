// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2020-07-13T11:26:46-07:00

package awsutils

// InstanceENIsAvailable contains a mapping of instance types to the number of ENIs available which is described at
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html#AvailableIpPerENI
var InstanceENIsAvailable = map[string]int{
	"a1.2xlarge":    4,
	"a1.4xlarge":    8,
	"a1.large":      3,
	"a1.medium":     2,
	"a1.metal":      8,
	"a1.xlarge":     4,
	"c1.medium":     2,
	"c1.xlarge":     4,
	"c3.2xlarge":    4,
	"c3.4xlarge":    8,
	"c3.8xlarge":    8,
	"c3.large":      3,
	"c3.xlarge":     4,
	"c4.2xlarge":    4,
	"c4.4xlarge":    8,
	"c4.8xlarge":    8,
	"c4.large":      3,
	"c4.xlarge":     4,
	"c5.12xlarge":   8,
	"c5.18xlarge":   15,
	"c5.24xlarge":   15,
	"c5.2xlarge":    4,
	"c5.4xlarge":    8,
	"c5.9xlarge":    8,
	"c5.large":      3,
	"c5.metal":      15,
	"c5.xlarge":     4,
	"c5a.12xlarge":  8,
	"c5a.16xlarge":  15,
	"c5a.24xlarge":  15,
	"c5a.2xlarge":   4,
	"c5a.4xlarge":   8,
	"c5a.8xlarge":   8,
	"c5a.large":     3,
	"c5a.metal":     15,
	"c5a.xlarge":    4,
	"c5ad.12xlarge": 8,
	"c5ad.16xlarge": 15,
	"c5ad.24xlarge": 15,
	"c5ad.2xlarge":  4,
	"c5ad.4xlarge":  8,
	"c5ad.8xlarge":  8,
	"c5ad.large":    3,
	"c5ad.metal":    15,
	"c5ad.xlarge":   4,
	"c5d.12xlarge":  8,
	"c5d.18xlarge":  15,
	"c5d.24xlarge":  15,
	"c5d.2xlarge":   4,
	"c5d.4xlarge":   8,
	"c5d.9xlarge":   8,
	"c5d.large":     3,
	"c5d.metal":     15,
	"c5d.xlarge":    4,
	"c5n.18xlarge":  15,
	"c5n.2xlarge":   4,
	"c5n.4xlarge":   8,
	"c5n.9xlarge":   8,
	"c5n.large":     3,
	"c5n.metal":     15,
	"c5n.xlarge":    4,
	"c6g.12xlarge":  8,
	"c6g.16xlarge":  15,
	"c6g.2xlarge":   4,
	"c6g.4xlarge":   8,
	"c6g.8xlarge":   8,
	"c6g.large":     3,
	"c6g.medium":    2,
	"c6g.metal":     15,
	"c6g.xlarge":    4,
	"cc2.8xlarge":   8,
	"cr1.8xlarge":   8,
	"d2.2xlarge":    4,
	"d2.4xlarge":    8,
	"d2.8xlarge":    8,
	"d2.xlarge":     4,
	"f1.16xlarge":   8,
	"f1.2xlarge":    4,
	"f1.4xlarge":    8,
	"g2.2xlarge":    4,
	"g2.8xlarge":    8,
	"g3.16xlarge":   15,
	"g3.4xlarge":    8,
	"g3.8xlarge":    8,
	"g3s.xlarge":    4,
	"g4dn.12xlarge": 8,
	"g4dn.16xlarge": 4,
	"g4dn.2xlarge":  3,
	"g4dn.4xlarge":  3,
	"g4dn.8xlarge":  4,
	"g4dn.metal":    15,
	"g4dn.xlarge":   3,
	"h1.16xlarge":   15,
	"h1.2xlarge":    4,
	"h1.4xlarge":    8,
	"h1.8xlarge":    8,
	"hs1.8xlarge":   8,
	"i2.2xlarge":    4,
	"i2.4xlarge":    8,
	"i2.8xlarge":    8,
	"i2.xlarge":     4,
	"i3.16xlarge":   15,
	"i3.2xlarge":    4,
	"i3.4xlarge":    8,
	"i3.8xlarge":    8,
	"i3.large":      3,
	"i3.metal":      15,
	"i3.xlarge":     4,
	"i3en.12xlarge": 8,
	"i3en.24xlarge": 15,
	"i3en.2xlarge":  4,
	"i3en.3xlarge":  4,
	"i3en.6xlarge":  8,
	"i3en.large":    3,
	"i3en.metal":    15,
	"i3en.xlarge":   4,
	"inf1.24xlarge": 15,
	"inf1.2xlarge":  4,
	"inf1.6xlarge":  8,
	"inf1.xlarge":   4,
	"m1.large":      3,
	"m1.medium":     2,
	"m1.small":      2,
	"m1.xlarge":     4,
	"m2.2xlarge":    4,
	"m2.4xlarge":    8,
	"m2.xlarge":     4,
	"m3.2xlarge":    4,
	"m3.large":      3,
	"m3.medium":     2,
	"m3.xlarge":     4,
	"m4.10xlarge":   8,
	"m4.16xlarge":   8,
	"m4.2xlarge":    4,
	"m4.4xlarge":    8,
	"m4.large":      2,
	"m4.xlarge":     4,
	"m5.12xlarge":   8,
	"m5.16xlarge":   15,
	"m5.24xlarge":   15,
	"m5.2xlarge":    4,
	"m5.4xlarge":    8,
	"m5.8xlarge":    8,
	"m5.large":      3,
	"m5.metal":      15,
	"m5.xlarge":     4,
	"m5a.12xlarge":  8,
	"m5a.16xlarge":  15,
	"m5a.24xlarge":  15,
	"m5a.2xlarge":   4,
	"m5a.4xlarge":   8,
	"m5a.8xlarge":   8,
	"m5a.large":     3,
	"m5a.xlarge":    4,
	"m5ad.12xlarge": 8,
	"m5ad.16xlarge": 15,
	"m5ad.24xlarge": 15,
	"m5ad.2xlarge":  4,
	"m5ad.4xlarge":  8,
	"m5ad.8xlarge":  8,
	"m5ad.large":    3,
	"m5ad.xlarge":   4,
	"m5d.12xlarge":  8,
	"m5d.16xlarge":  15,
	"m5d.24xlarge":  15,
	"m5d.2xlarge":   4,
	"m5d.4xlarge":   8,
	"m5d.8xlarge":   8,
	"m5d.large":     3,
	"m5d.metal":     15,
	"m5d.xlarge":    4,
	"m5dn.12xlarge": 8,
	"m5dn.16xlarge": 15,
	"m5dn.24xlarge": 15,
	"m5dn.2xlarge":  4,
	"m5dn.4xlarge":  8,
	"m5dn.8xlarge":  8,
	"m5dn.large":    3,
	"m5dn.xlarge":   4,
	"m5n.12xlarge":  8,
	"m5n.16xlarge":  15,
	"m5n.24xlarge":  15,
	"m5n.2xlarge":   4,
	"m5n.4xlarge":   8,
	"m5n.8xlarge":   8,
	"m5n.large":     3,
	"m5n.xlarge":    4,
	"m6g.12xlarge":  8,
	"m6g.16xlarge":  15,
	"m6g.2xlarge":   4,
	"m6g.4xlarge":   8,
	"m6g.8xlarge":   8,
	"m6g.large":     3,
	"m6g.medium":    2,
	"m6g.metal":     15,
	"m6g.xlarge":    4,
	"p2.16xlarge":   8,
	"p2.8xlarge":    8,
	"p2.xlarge":     4,
	"p3.16xlarge":   8,
	"p3.2xlarge":    4,
	"p3.8xlarge":    8,
	"p3dn.24xlarge": 15,
	"r3.2xlarge":    4,
	"r3.4xlarge":    8,
	"r3.8xlarge":    8,
	"r3.large":      3,
	"r3.xlarge":     4,
	"r4.16xlarge":   15,
	"r4.2xlarge":    4,
	"r4.4xlarge":    8,
	"r4.8xlarge":    8,
	"r4.large":      3,
	"r4.xlarge":     4,
	"r5.12xlarge":   8,
	"r5.16xlarge":   15,
	"r5.24xlarge":   15,
	"r5.2xlarge":    4,
	"r5.4xlarge":    8,
	"r5.8xlarge":    8,
	"r5.large":      3,
	"r5.metal":      15,
	"r5.xlarge":     4,
	"r5a.12xlarge":  8,
	"r5a.16xlarge":  15,
	"r5a.24xlarge":  15,
	"r5a.2xlarge":   4,
	"r5a.4xlarge":   8,
	"r5a.8xlarge":   8,
	"r5a.large":     3,
	"r5a.xlarge":    4,
	"r5ad.12xlarge": 8,
	"r5ad.16xlarge": 15,
	"r5ad.24xlarge": 15,
	"r5ad.2xlarge":  4,
	"r5ad.4xlarge":  8,
	"r5ad.8xlarge":  8,
	"r5ad.large":    3,
	"r5ad.xlarge":   4,
	"r5d.12xlarge":  8,
	"r5d.16xlarge":  15,
	"r5d.24xlarge":  15,
	"r5d.2xlarge":   4,
	"r5d.4xlarge":   8,
	"r5d.8xlarge":   8,
	"r5d.large":     3,
	"r5d.metal":     15,
	"r5d.xlarge":    4,
	"r5dn.12xlarge": 8,
	"r5dn.16xlarge": 15,
	"r5dn.24xlarge": 15,
	"r5dn.2xlarge":  4,
	"r5dn.4xlarge":  8,
	"r5dn.8xlarge":  8,
	"r5dn.large":    3,
	"r5dn.xlarge":   4,
	"r5n.12xlarge":  8,
	"r5n.16xlarge":  15,
	"r5n.24xlarge":  15,
	"r5n.2xlarge":   4,
	"r5n.4xlarge":   8,
	"r5n.8xlarge":   8,
	"r5n.large":     3,
	"r5n.xlarge":    4,
	"r6g.12xlarge":  8,
	"r6g.16xlarge":  15,
	"r6g.2xlarge":   4,
	"r6g.4xlarge":   8,
	"r6g.8xlarge":   8,
	"r6g.large":     3,
	"r6g.medium":    2,
	"r6g.metal":     15,
	"r6g.xlarge":    4,
	"t1.micro":      2,
	"t2.2xlarge":    3,
	"t2.large":      3,
	"t2.medium":     3,
	"t2.micro":      2,
	"t2.nano":       2,
	"t2.small":      3,
	"t2.xlarge":     3,
	"t3.2xlarge":    4,
	"t3.large":      3,
	"t3.medium":     3,
	"t3.micro":      2,
	"t3.nano":       2,
	"t3.small":      3,
	"t3.xlarge":     4,
	"t3a.2xlarge":   4,
	"t3a.large":     3,
	"t3a.medium":    3,
	"t3a.micro":     2,
	"t3a.nano":      2,
	"t3a.small":     2,
	"t3a.xlarge":    4,
	"u-12tb1.metal": 5,
	"u-18tb1.metal": 15,
	"u-24tb1.metal": 15,
	"u-6tb1.metal":  5,
	"u-9tb1.metal":  5,
	"x1.16xlarge":   8,
	"x1.32xlarge":   8,
	"x1e.16xlarge":  8,
	"x1e.2xlarge":   4,
	"x1e.32xlarge":  8,
	"x1e.4xlarge":   4,
	"x1e.8xlarge":   4,
	"x1e.xlarge":    3,
	"z1d.12xlarge":  15,
	"z1d.2xlarge":   4,
	"z1d.3xlarge":   8,
	"z1d.6xlarge":   8,
	"z1d.large":     3,
	"z1d.metal":     15,
	"z1d.xlarge":    4,
}

// InstanceIPsAvailable contains a mapping of instance types to the number of IPs per ENI
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html#AvailableIpPerENI
var InstanceIPsAvailable = map[string]int{
	"a1.2xlarge":    15,
	"a1.4xlarge":    30,
	"a1.large":      10,
	"a1.medium":     4,
	"a1.metal":      30,
	"a1.xlarge":     15,
	"c1.medium":     6,
	"c1.xlarge":     15,
	"c3.2xlarge":    15,
	"c3.4xlarge":    30,
	"c3.8xlarge":    30,
	"c3.large":      10,
	"c3.xlarge":     15,
	"c4.2xlarge":    15,
	"c4.4xlarge":    30,
	"c4.8xlarge":    30,
	"c4.large":      10,
	"c4.xlarge":     15,
	"c5.12xlarge":   30,
	"c5.18xlarge":   50,
	"c5.24xlarge":   50,
	"c5.2xlarge":    15,
	"c5.4xlarge":    30,
	"c5.9xlarge":    30,
	"c5.large":      10,
	"c5.metal":      50,
	"c5.xlarge":     15,
	"c5a.12xlarge":  30,
	"c5a.16xlarge":  50,
	"c5a.24xlarge":  50,
	"c5a.2xlarge":   15,
	"c5a.4xlarge":   30,
	"c5a.8xlarge":   30,
	"c5a.large":     10,
	"c5a.metal":     50,
	"c5a.xlarge":    15,
	"c5ad.12xlarge": 30,
	"c5ad.16xlarge": 50,
	"c5ad.24xlarge": 50,
	"c5ad.2xlarge":  15,
	"c5ad.4xlarge":  30,
	"c5ad.8xlarge":  30,
	"c5ad.large":    10,
	"c5ad.metal":    50,
	"c5ad.xlarge":   15,
	"c5d.12xlarge":  30,
	"c5d.18xlarge":  50,
	"c5d.24xlarge":  50,
	"c5d.2xlarge":   15,
	"c5d.4xlarge":   30,
	"c5d.9xlarge":   30,
	"c5d.large":     10,
	"c5d.metal":     50,
	"c5d.xlarge":    15,
	"c5n.18xlarge":  50,
	"c5n.2xlarge":   15,
	"c5n.4xlarge":   30,
	"c5n.9xlarge":   30,
	"c5n.large":     10,
	"c5n.metal":     50,
	"c5n.xlarge":    15,
	"c6g.12xlarge":  30,
	"c6g.16xlarge":  50,
	"c6g.2xlarge":   15,
	"c6g.4xlarge":   30,
	"c6g.8xlarge":   30,
	"c6g.large":     10,
	"c6g.medium":    4,
	"c6g.metal":     50,
	"c6g.xlarge":    15,
	"cc2.8xlarge":   30,
	"cr1.8xlarge":   30,
	"d2.2xlarge":    15,
	"d2.4xlarge":    30,
	"d2.8xlarge":    30,
	"d2.xlarge":     15,
	"f1.16xlarge":   50,
	"f1.2xlarge":    15,
	"f1.4xlarge":    30,
	"g2.2xlarge":    15,
	"g2.8xlarge":    30,
	"g3.16xlarge":   50,
	"g3.4xlarge":    30,
	"g3.8xlarge":    30,
	"g3s.xlarge":    15,
	"g4dn.12xlarge": 30,
	"g4dn.16xlarge": 15,
	"g4dn.2xlarge":  10,
	"g4dn.4xlarge":  10,
	"g4dn.8xlarge":  15,
	"g4dn.metal":    50,
	"g4dn.xlarge":   10,
	"h1.16xlarge":   50,
	"h1.2xlarge":    15,
	"h1.4xlarge":    30,
	"h1.8xlarge":    30,
	"hs1.8xlarge":   30,
	"i2.2xlarge":    15,
	"i2.4xlarge":    30,
	"i2.8xlarge":    30,
	"i2.xlarge":     15,
	"i3.16xlarge":   50,
	"i3.2xlarge":    15,
	"i3.4xlarge":    30,
	"i3.8xlarge":    30,
	"i3.large":      10,
	"i3.metal":      50,
	"i3.xlarge":     15,
	"i3en.12xlarge": 30,
	"i3en.24xlarge": 50,
	"i3en.2xlarge":  15,
	"i3en.3xlarge":  15,
	"i3en.6xlarge":  30,
	"i3en.large":    10,
	"i3en.metal":    50,
	"i3en.xlarge":   15,
	"inf1.24xlarge": 30,
	"inf1.2xlarge":  10,
	"inf1.6xlarge":  30,
	"inf1.xlarge":   10,
	"m1.large":      10,
	"m1.medium":     6,
	"m1.small":      4,
	"m1.xlarge":     15,
	"m2.2xlarge":    30,
	"m2.4xlarge":    30,
	"m2.xlarge":     15,
	"m3.2xlarge":    30,
	"m3.large":      10,
	"m3.medium":     6,
	"m3.xlarge":     15,
	"m4.10xlarge":   30,
	"m4.16xlarge":   30,
	"m4.2xlarge":    15,
	"m4.4xlarge":    30,
	"m4.large":      10,
	"m4.xlarge":     15,
	"m5.12xlarge":   30,
	"m5.16xlarge":   50,
	"m5.24xlarge":   50,
	"m5.2xlarge":    15,
	"m5.4xlarge":    30,
	"m5.8xlarge":    30,
	"m5.large":      10,
	"m5.metal":      50,
	"m5.xlarge":     15,
	"m5a.12xlarge":  30,
	"m5a.16xlarge":  50,
	"m5a.24xlarge":  50,
	"m5a.2xlarge":   15,
	"m5a.4xlarge":   30,
	"m5a.8xlarge":   30,
	"m5a.large":     10,
	"m5a.xlarge":    15,
	"m5ad.12xlarge": 30,
	"m5ad.16xlarge": 50,
	"m5ad.24xlarge": 50,
	"m5ad.2xlarge":  15,
	"m5ad.4xlarge":  30,
	"m5ad.8xlarge":  30,
	"m5ad.large":    10,
	"m5ad.xlarge":   15,
	"m5d.12xlarge":  30,
	"m5d.16xlarge":  50,
	"m5d.24xlarge":  50,
	"m5d.2xlarge":   15,
	"m5d.4xlarge":   30,
	"m5d.8xlarge":   30,
	"m5d.large":     10,
	"m5d.metal":     50,
	"m5d.xlarge":    15,
	"m5dn.12xlarge": 30,
	"m5dn.16xlarge": 50,
	"m5dn.24xlarge": 50,
	"m5dn.2xlarge":  15,
	"m5dn.4xlarge":  30,
	"m5dn.8xlarge":  30,
	"m5dn.large":    10,
	"m5dn.xlarge":   15,
	"m5n.12xlarge":  30,
	"m5n.16xlarge":  50,
	"m5n.24xlarge":  50,
	"m5n.2xlarge":   15,
	"m5n.4xlarge":   30,
	"m5n.8xlarge":   30,
	"m5n.large":     10,
	"m5n.xlarge":    15,
	"m6g.12xlarge":  30,
	"m6g.16xlarge":  50,
	"m6g.2xlarge":   15,
	"m6g.4xlarge":   30,
	"m6g.8xlarge":   30,
	"m6g.large":     10,
	"m6g.medium":    4,
	"m6g.metal":     50,
	"m6g.xlarge":    15,
	"p2.16xlarge":   30,
	"p2.8xlarge":    30,
	"p2.xlarge":     15,
	"p3.16xlarge":   30,
	"p3.2xlarge":    15,
	"p3.8xlarge":    30,
	"p3dn.24xlarge": 50,
	"r3.2xlarge":    15,
	"r3.4xlarge":    30,
	"r3.8xlarge":    30,
	"r3.large":      10,
	"r3.xlarge":     15,
	"r4.16xlarge":   50,
	"r4.2xlarge":    15,
	"r4.4xlarge":    30,
	"r4.8xlarge":    30,
	"r4.large":      10,
	"r4.xlarge":     15,
	"r5.12xlarge":   30,
	"r5.16xlarge":   50,
	"r5.24xlarge":   50,
	"r5.2xlarge":    15,
	"r5.4xlarge":    30,
	"r5.8xlarge":    30,
	"r5.large":      10,
	"r5.metal":      50,
	"r5.xlarge":     15,
	"r5a.12xlarge":  30,
	"r5a.16xlarge":  50,
	"r5a.24xlarge":  50,
	"r5a.2xlarge":   15,
	"r5a.4xlarge":   30,
	"r5a.8xlarge":   30,
	"r5a.large":     10,
	"r5a.xlarge":    15,
	"r5ad.12xlarge": 30,
	"r5ad.16xlarge": 50,
	"r5ad.24xlarge": 50,
	"r5ad.2xlarge":  15,
	"r5ad.4xlarge":  30,
	"r5ad.8xlarge":  30,
	"r5ad.large":    10,
	"r5ad.xlarge":   15,
	"r5d.12xlarge":  30,
	"r5d.16xlarge":  50,
	"r5d.24xlarge":  50,
	"r5d.2xlarge":   15,
	"r5d.4xlarge":   30,
	"r5d.8xlarge":   30,
	"r5d.large":     10,
	"r5d.metal":     50,
	"r5d.xlarge":    15,
	"r5dn.12xlarge": 30,
	"r5dn.16xlarge": 50,
	"r5dn.24xlarge": 50,
	"r5dn.2xlarge":  15,
	"r5dn.4xlarge":  30,
	"r5dn.8xlarge":  30,
	"r5dn.large":    10,
	"r5dn.xlarge":   15,
	"r5n.12xlarge":  30,
	"r5n.16xlarge":  50,
	"r5n.24xlarge":  50,
	"r5n.2xlarge":   15,
	"r5n.4xlarge":   30,
	"r5n.8xlarge":   30,
	"r5n.large":     10,
	"r5n.xlarge":    15,
	"r6g.12xlarge":  30,
	"r6g.16xlarge":  50,
	"r6g.2xlarge":   15,
	"r6g.4xlarge":   30,
	"r6g.8xlarge":   30,
	"r6g.large":     10,
	"r6g.medium":    4,
	"r6g.metal":     50,
	"r6g.xlarge":    15,
	"t1.micro":      2,
	"t2.2xlarge":    15,
	"t2.large":      12,
	"t2.medium":     6,
	"t2.micro":      2,
	"t2.nano":       2,
	"t2.small":      4,
	"t2.xlarge":     15,
	"t3.2xlarge":    15,
	"t3.large":      12,
	"t3.medium":     6,
	"t3.micro":      2,
	"t3.nano":       2,
	"t3.small":      4,
	"t3.xlarge":     15,
	"t3a.2xlarge":   15,
	"t3a.large":     12,
	"t3a.medium":    6,
	"t3a.micro":     2,
	"t3a.nano":      2,
	"t3a.small":     4,
	"t3a.xlarge":    15,
	"u-12tb1.metal": 30,
	"u-18tb1.metal": 50,
	"u-24tb1.metal": 50,
	"u-6tb1.metal":  30,
	"u-9tb1.metal":  30,
	"x1.16xlarge":   30,
	"x1.32xlarge":   30,
	"x1e.16xlarge":  30,
	"x1e.2xlarge":   15,
	"x1e.32xlarge":  30,
	"x1e.4xlarge":   15,
	"x1e.8xlarge":   15,
	"x1e.xlarge":    10,
	"z1d.12xlarge":  50,
	"z1d.2xlarge":   15,
	"z1d.3xlarge":   30,
	"z1d.6xlarge":   30,
	"z1d.large":     10,
	"z1d.metal":     50,
	"z1d.xlarge":    15,
}
