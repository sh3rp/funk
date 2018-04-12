package io

import "github.com/sh3rp/funk/funk"

type Input interface {
	Next() funk.Event
}

type Output interface {
	Publish(funk.Event)
}
