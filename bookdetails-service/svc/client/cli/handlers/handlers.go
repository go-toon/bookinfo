// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 7dc4d5d85c
// Version Date: 2018年 5月28日 星期一 22时12分59秒 UTC

package handlers

import (
	pb "bookinfo/pb/details"
)

// Detail implements Service.
func Detail(IdDetail int64) (*pb.DetailReq, error) {
	request := pb.DetailReq{
		Id: IdDetail,
	}
	return &request, nil
}
