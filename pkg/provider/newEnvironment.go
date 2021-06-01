// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"github.com/pulumi/pulumi-gcp/v5/sdk/go/gcp/composer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnvironmentArgs struct {
	Region: pulumi.String("us-central1")
	Config EnvironmentConfigOutput `pulumi:"config"`
}

type Environment struct {
	pulumi.ResourceState

	// Bucket     *s3.Bucket          `pulumi:"bucket"`
	// WebsiteUrl pulumi.StringOutput `pulumi:"websiteUrl"`
}

func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil {
		args = &EnvironmentArgs{}
	}

	component := &Environment{}
	err := ctx.RegisterComponentResource("gcp:composer/environment:Environment", name, component, opts...)
	if err != nil {
		return nil, err
	}

	return component, nil
}
