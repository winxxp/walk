// THIS FILE WAS GENERATED BY A TOOL, DO NOT EDIT!

package main

import (
	"github.com/winxxp/walk"
)

type myDialogUI struct {
	noProgressBtn    *walk.PushButton
	indeterminateBtn *walk.PushButton
	normalBtn        *walk.PushButton
	errBtn           *walk.PushButton
	pausedBtn        *walk.PushButton
	startBtn         *walk.PushButton
}

func (w *MyDialog) init(owner walk.Form) (err error) {
	if w.Dialog, err = walk.NewDialog(owner); err != nil {
		return err
	}

	succeeded := false
	defer func() {
		if !succeeded {
			w.Dispose()
		}
	}()

	var font *walk.Font
	if font == nil {
		font = nil
	}

	w.SetName("Dialog")
	if err := w.SetClientSize(walk.Size{598, 300}); err != nil {
		return err
	}
	if err := w.SetTitle(`Dialog`); err != nil {
		return err
	}

	// noProgressBtn
	if w.ui.noProgressBtn, err = walk.NewPushButton(w); err != nil {
		return err
	}
	w.ui.noProgressBtn.SetName("noProgressBtn")
	if err := w.ui.noProgressBtn.SetBounds(walk.Rectangle{40, 60, 161, 23}); err != nil {
		return err
	}
	if err := w.ui.noProgressBtn.SetText(`NoProgress`); err != nil {
		return err
	}
	if err := w.ui.noProgressBtn.SetMinMaxSize(walk.Size{0, 0}, walk.Size{161, 16777215}); err != nil {
		return err
	}

	// indeterminateBtn
	if w.ui.indeterminateBtn, err = walk.NewPushButton(w); err != nil {
		return err
	}
	w.ui.indeterminateBtn.SetName("indeterminateBtn")
	if err := w.ui.indeterminateBtn.SetBounds(walk.Rectangle{40, 90, 161, 23}); err != nil {
		return err
	}
	if err := w.ui.indeterminateBtn.SetText(`Indeterminate`); err != nil {
		return err
	}
	if err := w.ui.indeterminateBtn.SetMinMaxSize(walk.Size{0, 0}, walk.Size{161, 16777215}); err != nil {
		return err
	}

	// normalBtn
	if w.ui.normalBtn, err = walk.NewPushButton(w); err != nil {
		return err
	}
	w.ui.normalBtn.SetName("normalBtn")
	if err := w.ui.normalBtn.SetBounds(walk.Rectangle{40, 120, 161, 23}); err != nil {
		return err
	}
	if err := w.ui.normalBtn.SetText(`Normal`); err != nil {
		return err
	}
	if err := w.ui.normalBtn.SetMinMaxSize(walk.Size{0, 0}, walk.Size{161, 16777215}); err != nil {
		return err
	}

	// errBtn
	if w.ui.errBtn, err = walk.NewPushButton(w); err != nil {
		return err
	}
	w.ui.errBtn.SetName("errBtn")
	if err := w.ui.errBtn.SetBounds(walk.Rectangle{40, 150, 161, 23}); err != nil {
		return err
	}
	if err := w.ui.errBtn.SetText(`Error`); err != nil {
		return err
	}

	// pausedBtn
	if w.ui.pausedBtn, err = walk.NewPushButton(w); err != nil {
		return err
	}
	w.ui.pausedBtn.SetName("pausedBtn")
	if err := w.ui.pausedBtn.SetBounds(walk.Rectangle{40, 180, 161, 23}); err != nil {
		return err
	}
	if err := w.ui.pausedBtn.SetText(`Paused`); err != nil {
		return err
	}

	// startBtn
	if w.ui.startBtn, err = walk.NewPushButton(w); err != nil {
		return err
	}
	w.ui.startBtn.SetName("startBtn")
	if err := w.ui.startBtn.SetBounds(walk.Rectangle{290, 180, 75, 23}); err != nil {
		return err
	}
	if err := w.ui.startBtn.SetText(`START`); err != nil {
		return err
	}

	// Tab order

	succeeded = true

	return nil
}
