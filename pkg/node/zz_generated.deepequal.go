//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by deepequal-gen. DO NOT EDIT.

package node

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *LocalNode) DeepEqual(other *LocalNode) bool {
	if other == nil {
		return false
	}

	if !in.Node.DeepEqual(&other.Node) {
		return false
	}

	if in.OptOutNodeEncryption != other.OptOutNodeEncryption {
		return false
	}
	if in.UID != other.UID {
		return false
	}
	if in.ProviderID != other.ProviderID {
		return false
	}
	if (in.IPv4NativeRoutingCIDR == nil) != (other.IPv4NativeRoutingCIDR == nil) {
		return false
	} else if in.IPv4NativeRoutingCIDR != nil {
		if !in.IPv4NativeRoutingCIDR.DeepEqual(other.IPv4NativeRoutingCIDR) {
			return false
		}
	}

	if (in.IPv6NativeRoutingCIDR == nil) != (other.IPv6NativeRoutingCIDR == nil) {
		return false
	} else if in.IPv6NativeRoutingCIDR != nil {
		if !in.IPv6NativeRoutingCIDR.DeepEqual(other.IPv6NativeRoutingCIDR) {
			return false
		}
	}

	if ((in.IPv4Loopback != nil) && (other.IPv4Loopback != nil)) || ((in.IPv4Loopback == nil) != (other.IPv4Loopback == nil)) {
		in, other := &in.IPv4Loopback, &other.IPv4Loopback
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	return true
}