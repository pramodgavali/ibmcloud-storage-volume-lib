/*******************************************************************************
 * IBM Confidential
 * OCO Source Materials
 * IBM Cloud Container Service, 5737-D43
 * (C) Copyright IBM Corp. 2018 All Rights Reserved.
 * The source code for this program is not  published or otherwise divested of
 * its trade secrets, irrespective of what has been deposited with
 * the U.S. Copyright Office.
 ******************************************************************************/

package models

//GPU ...
type GPU struct {
	Cores        int64  `json:"cores,omitempty"`
	Count        int64  `json:"count,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	Memory       int64  `json:"memory,omitempty"`
	Model        string `json:"model,omitempty"`
}