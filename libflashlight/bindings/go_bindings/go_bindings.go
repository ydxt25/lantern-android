// Package go_flashlight is an autogenerated binder stub for package flashlight.
//   gobind -lang=go github.com/getlantern/lantern-android/libflashlight/bindings
//
// File is generated by gobind. Do not edit.
package go_flashlight

import (
	"github.com/getlantern/lantern-android/libflashlight/bindings"
	"golang.org/x/mobile/bind/seq"
)

func proxy_RunClientProxy(out, in *seq.Buffer) {
	param_listenAddr := in.ReadUTF16()
	err := flashlight.RunClientProxy(param_listenAddr)
	if err == nil {
		out.WriteUTF16("")
	} else {
		out.WriteUTF16(err.Error())
	}
}

func proxy_StopClientProxy(out, in *seq.Buffer) {
	err := flashlight.StopClientProxy()
	if err == nil {
		out.WriteUTF16("")
	} else {
		out.WriteUTF16(err.Error())
	}
}

func init() {
	seq.Register("flashlight", 1, proxy_RunClientProxy)
	seq.Register("flashlight", 2, proxy_StopClientProxy)
}
