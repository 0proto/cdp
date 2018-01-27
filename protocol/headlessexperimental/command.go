// Code generated by cdpgen. DO NOT EDIT.

package headlessexperimental

import (
	"github.com/mafredri/cdp/protocol/runtime"
)

// BeginFrameArgs represents the arguments for BeginFrame in the HeadlessExperimental domain.
type BeginFrameArgs struct {
	FrameTime        *runtime.Timestamp `json:"frameTime,omitempty"`        // Timestamp of this BeginFrame (milliseconds since epoch). If not set, the current time will be used.
	Deadline         *runtime.Timestamp `json:"deadline,omitempty"`         // Deadline of this BeginFrame (milliseconds since epoch). If not set, the deadline will be calculated from the frameTime and interval.
	Interval         *float64           `json:"interval,omitempty"`         // The interval between BeginFrames that is reported to the compositor, in milliseconds. Defaults to a 60 frames/second interval, i.e. about 16.666 milliseconds.
	NoDisplayUpdates *bool              `json:"noDisplayUpdates,omitempty"` // Whether updates should not be committed and drawn onto the display. False by default. If true, only side effects of the BeginFrame will be run, such as layout and animations, but any visual updates may not be visible on the display or in screenshots.
	Screenshot       *ScreenshotParams  `json:"screenshot,omitempty"`       // If set, a screenshot of the frame will be captured and returned in the response. Otherwise, no screenshot will be captured.
}

// NewBeginFrameArgs initializes BeginFrameArgs with the required arguments.
func NewBeginFrameArgs() *BeginFrameArgs {
	args := new(BeginFrameArgs)

	return args
}

// SetFrameTime sets the FrameTime optional argument. Timestamp of
// this BeginFrame (milliseconds since epoch). If not set, the current
// time will be used.
func (a *BeginFrameArgs) SetFrameTime(frameTime runtime.Timestamp) *BeginFrameArgs {
	a.FrameTime = &frameTime
	return a
}

// SetDeadline sets the Deadline optional argument. Deadline of this
// BeginFrame (milliseconds since epoch). If not set, the deadline will
// be calculated from the frameTime and interval.
func (a *BeginFrameArgs) SetDeadline(deadline runtime.Timestamp) *BeginFrameArgs {
	a.Deadline = &deadline
	return a
}

// SetInterval sets the Interval optional argument. The interval
// between BeginFrames that is reported to the compositor, in
// milliseconds. Defaults to a 60 frames/second interval, i.e. about
// 16.666 milliseconds.
func (a *BeginFrameArgs) SetInterval(interval float64) *BeginFrameArgs {
	a.Interval = &interval
	return a
}

// SetNoDisplayUpdates sets the NoDisplayUpdates optional argument.
// Whether updates should not be committed and drawn onto the display.
// False by default. If true, only side effects of the BeginFrame will
// be run, such as layout and animations, but any visual updates may
// not be visible on the display or in screenshots.
func (a *BeginFrameArgs) SetNoDisplayUpdates(noDisplayUpdates bool) *BeginFrameArgs {
	a.NoDisplayUpdates = &noDisplayUpdates
	return a
}

// SetScreenshot sets the Screenshot optional argument. If set, a
// screenshot of the frame will be captured and returned in the
// response. Otherwise, no screenshot will be captured.
func (a *BeginFrameArgs) SetScreenshot(screenshot ScreenshotParams) *BeginFrameArgs {
	a.Screenshot = &screenshot
	return a
}

// BeginFrameReply represents the return values for BeginFrame in the HeadlessExperimental domain.
type BeginFrameReply struct {
	HasDamage               bool   `json:"hasDamage"`                // Whether the BeginFrame resulted in damage and, thus, a new frame was committed to the display.
	MainFrameContentUpdated bool   `json:"mainFrameContentUpdated"`  // Whether the main frame submitted a new display frame in response to this BeginFrame.
	ScreenshotData          []byte `json:"screenshotData,omitempty"` // Base64-encoded image data of the screenshot, if one was requested and successfully taken.
}
