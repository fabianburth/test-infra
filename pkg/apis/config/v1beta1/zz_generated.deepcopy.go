//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BotConfiguration) DeepCopyInto(out *BotConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Webserver = in.Webserver
	in.Dashboard.DeepCopyInto(&out.Dashboard)
	out.GitHubBot = in.GitHubBot
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BotConfiguration.
func (in *BotConfiguration) DeepCopy() *BotConfiguration {
	if in == nil {
		return nil
	}
	out := new(BotConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BotConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certificate) DeepCopyInto(out *Certificate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certificate.
func (in *Certificate) DeepCopy() *Certificate {
	if in == nil {
		return nil
	}
	out := new(Certificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Controller = in.Controller
	in.TestMachinery.DeepCopyInto(&out.TestMachinery)
	in.GitHub.DeepCopyInto(&out.GitHub)
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3)
		**out = **in
	}
	if in.ElasticSearch != nil {
		in, out := &in.ElasticSearch, &out.ElasticSearch
		*out = new(ElasticSearch)
		**out = **in
	}
	if in.ImagePullSecretNames != nil {
		in, out := &in.ImagePullSecretNames, &out.ImagePullSecretNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Configuration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Controller) DeepCopyInto(out *Controller) {
	*out = *in
	out.TTLController = in.TTLController
	out.WebhookConfig = in.WebhookConfig
	out.DependencyHealthCheck = in.DependencyHealthCheck
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Controller.
func (in *Controller) DeepCopy() *Controller {
	if in == nil {
		return nil
	}
	out := new(Controller)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dashboard) DeepCopyInto(out *Dashboard) {
	*out = *in
	in.Authentication.DeepCopyInto(&out.Authentication)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dashboard.
func (in *Dashboard) DeepCopy() *Dashboard {
	if in == nil {
		return nil
	}
	out := new(Dashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardAuthentication) DeepCopyInto(out *DashboardAuthentication) {
	*out = *in
	if in.GitHub != nil {
		in, out := &in.GitHub, &out.GitHub
		*out = new(GitHubAuthentication)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardAuthentication.
func (in *DashboardAuthentication) DeepCopy() *DashboardAuthentication {
	if in == nil {
		return nil
	}
	out := new(DashboardAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticSearch) DeepCopyInto(out *ElasticSearch) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticSearch.
func (in *ElasticSearch) DeepCopy() *ElasticSearch {
	if in == nil {
		return nil
	}
	out := new(ElasticSearch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHub) DeepCopyInto(out *GitHub) {
	*out = *in
	if in.Cache != nil {
		in, out := &in.Cache, &out.Cache
		*out = new(GitHubCache)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHub.
func (in *GitHub) DeepCopy() *GitHub {
	if in == nil {
		return nil
	}
	out := new(GitHub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHubAuthentication) DeepCopyInto(out *GitHubAuthentication) {
	*out = *in
	if in.OAuth != nil {
		in, out := &in.OAuth, &out.OAuth
		*out = new(OAuth)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHubAuthentication.
func (in *GitHubAuthentication) DeepCopy() *GitHubAuthentication {
	if in == nil {
		return nil
	}
	out := new(GitHubAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHubBot) DeepCopyInto(out *GitHubBot) {
	*out = *in
	out.GitHubCache = in.GitHubCache
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHubBot.
func (in *GitHubBot) DeepCopy() *GitHubBot {
	if in == nil {
		return nil
	}
	out := new(GitHubBot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHubCache) DeepCopyInto(out *GitHubCache) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHubCache.
func (in *GitHubCache) DeepCopy() *GitHubCache {
	if in == nil {
		return nil
	}
	out := new(GitHubCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckTarget) DeepCopyInto(out *HealthCheckTarget) {
	*out = *in
	out.Interval = in.Interval
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckTarget.
func (in *HealthCheckTarget) DeepCopy() *HealthCheckTarget {
	if in == nil {
		return nil
	}
	out := new(HealthCheckTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LandscapeMapping) DeepCopyInto(out *LandscapeMapping) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LandscapeMapping.
func (in *LandscapeMapping) DeepCopy() *LandscapeMapping {
	if in == nil {
		return nil
	}
	out := new(LandscapeMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Locations) DeepCopyInto(out *Locations) {
	*out = *in
	if in.ExcludeDomains != nil {
		in, out := &in.ExcludeDomains, &out.ExcludeDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Locations.
func (in *Locations) DeepCopy() *Locations {
	if in == nil {
		return nil
	}
	out := new(Locations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OAuth) DeepCopyInto(out *OAuth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OAuth.
func (in *OAuth) DeepCopy() *OAuth {
	if in == nil {
		return nil
	}
	out := new(OAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3) DeepCopyInto(out *S3) {
	*out = *in
	out.Server = in.Server
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3.
func (in *S3) DeepCopy() *S3 {
	if in == nil {
		return nil
	}
	out := new(S3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Server) DeepCopyInto(out *S3Server) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Server.
func (in *S3Server) DeepCopy() *S3Server {
	if in == nil {
		return nil
	}
	out := new(S3Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TTLController) DeepCopyInto(out *TTLController) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TTLController.
func (in *TTLController) DeepCopy() *TTLController {
	if in == nil {
		return nil
	}
	out := new(TTLController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestMachinery) DeepCopyInto(out *TestMachinery) {
	*out = *in
	in.Locations.DeepCopyInto(&out.Locations)
	if in.RetryTimeoutDuration != nil {
		in, out := &in.RetryTimeoutDuration, &out.RetryTimeoutDuration
		*out = new(time.Duration)
		**out = **in
	}
	if in.LandscapeMappings != nil {
		in, out := &in.LandscapeMappings, &out.LandscapeMappings
		*out = make([]LandscapeMapping, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestMachinery.
func (in *TestMachinery) DeepCopy() *TestMachinery {
	if in == nil {
		return nil
	}
	out := new(TestMachinery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConfig) DeepCopyInto(out *WebhookConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConfig.
func (in *WebhookConfig) DeepCopy() *WebhookConfig {
	if in == nil {
		return nil
	}
	out := new(WebhookConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webserver) DeepCopyInto(out *Webserver) {
	*out = *in
	out.Certificate = in.Certificate
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webserver.
func (in *Webserver) DeepCopy() *Webserver {
	if in == nil {
		return nil
	}
	out := new(Webserver)
	in.DeepCopyInto(out)
	return out
}
