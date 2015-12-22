package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QProxyStyle struct {
	QCommonStyle
}

type QProxyStyle_ITF interface {
	QCommonStyle_ITF
	QProxyStyle_PTR() *QProxyStyle
}

func PointerFromQProxyStyle(ptr QProxyStyle_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProxyStyle_PTR().Pointer()
	}
	return nil
}

func NewQProxyStyleFromPointer(ptr unsafe.Pointer) *QProxyStyle {
	var n = new(QProxyStyle)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QProxyStyle_") {
		n.SetObjectName("QProxyStyle_" + qt.Identifier())
	}
	return n
}

func (ptr *QProxyStyle) QProxyStyle_PTR() *QProxyStyle {
	return ptr
}

func (ptr *QProxyStyle) BaseStyle() *QStyle {
	defer qt.Recovering("QProxyStyle::baseStyle")

	if ptr.Pointer() != nil {
		return NewQStyleFromPointer(C.QProxyStyle_BaseStyle(ptr.Pointer()))
	}
	return nil
}

func (ptr *QProxyStyle) ConnectDrawComplexControl(f func(control QStyle__ComplexControl, option *QStyleOptionComplex, painter *gui.QPainter, widget *QWidget)) {
	defer qt.Recovering("connect QProxyStyle::drawComplexControl")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "drawComplexControl", f)
	}
}

func (ptr *QProxyStyle) DisconnectDrawComplexControl() {
	defer qt.Recovering("disconnect QProxyStyle::drawComplexControl")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "drawComplexControl")
	}
}

//export callbackQProxyStyleDrawComplexControl
func callbackQProxyStyleDrawComplexControl(ptrName *C.char, control C.int, option unsafe.Pointer, painter unsafe.Pointer, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QProxyStyle::drawComplexControl")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawComplexControl"); signal != nil {
		signal.(func(QStyle__ComplexControl, *QStyleOptionComplex, *gui.QPainter, *QWidget))(QStyle__ComplexControl(control), NewQStyleOptionComplexFromPointer(option), gui.NewQPainterFromPointer(painter), NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QProxyStyle) ConnectDrawControl(f func(element QStyle__ControlElement, option *QStyleOption, painter *gui.QPainter, widget *QWidget)) {
	defer qt.Recovering("connect QProxyStyle::drawControl")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "drawControl", f)
	}
}

func (ptr *QProxyStyle) DisconnectDrawControl() {
	defer qt.Recovering("disconnect QProxyStyle::drawControl")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "drawControl")
	}
}

//export callbackQProxyStyleDrawControl
func callbackQProxyStyleDrawControl(ptrName *C.char, element C.int, option unsafe.Pointer, painter unsafe.Pointer, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QProxyStyle::drawControl")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawControl"); signal != nil {
		signal.(func(QStyle__ControlElement, *QStyleOption, *gui.QPainter, *QWidget))(QStyle__ControlElement(element), NewQStyleOptionFromPointer(option), gui.NewQPainterFromPointer(painter), NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QProxyStyle) ConnectDrawPrimitive(f func(element QStyle__PrimitiveElement, option *QStyleOption, painter *gui.QPainter, widget *QWidget)) {
	defer qt.Recovering("connect QProxyStyle::drawPrimitive")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "drawPrimitive", f)
	}
}

func (ptr *QProxyStyle) DisconnectDrawPrimitive() {
	defer qt.Recovering("disconnect QProxyStyle::drawPrimitive")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "drawPrimitive")
	}
}

//export callbackQProxyStyleDrawPrimitive
func callbackQProxyStyleDrawPrimitive(ptrName *C.char, element C.int, option unsafe.Pointer, painter unsafe.Pointer, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QProxyStyle::drawPrimitive")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawPrimitive"); signal != nil {
		signal.(func(QStyle__PrimitiveElement, *QStyleOption, *gui.QPainter, *QWidget))(QStyle__PrimitiveElement(element), NewQStyleOptionFromPointer(option), gui.NewQPainterFromPointer(painter), NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QProxyStyle) HitTestComplexControl(control QStyle__ComplexControl, option QStyleOptionComplex_ITF, pos core.QPoint_ITF, widget QWidget_ITF) QStyle__SubControl {
	defer qt.Recovering("QProxyStyle::hitTestComplexControl")

	if ptr.Pointer() != nil {
		return QStyle__SubControl(C.QProxyStyle_HitTestComplexControl(ptr.Pointer(), C.int(control), PointerFromQStyleOptionComplex(option), core.PointerFromQPoint(pos), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QProxyStyle) ItemPixmapRect(r core.QRect_ITF, flags int, pixmap gui.QPixmap_ITF) *core.QRect {
	defer qt.Recovering("QProxyStyle::itemPixmapRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QProxyStyle_ItemPixmapRect(ptr.Pointer(), core.PointerFromQRect(r), C.int(flags), gui.PointerFromQPixmap(pixmap)))
	}
	return nil
}

func (ptr *QProxyStyle) ItemTextRect(fm gui.QFontMetrics_ITF, r core.QRect_ITF, flags int, enabled bool, text string) *core.QRect {
	defer qt.Recovering("QProxyStyle::itemTextRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QProxyStyle_ItemTextRect(ptr.Pointer(), gui.PointerFromQFontMetrics(fm), core.PointerFromQRect(r), C.int(flags), C.int(qt.GoBoolToInt(enabled)), C.CString(text)))
	}
	return nil
}

func (ptr *QProxyStyle) LayoutSpacing(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation core.Qt__Orientation, option QStyleOption_ITF, widget QWidget_ITF) int {
	defer qt.Recovering("QProxyStyle::layoutSpacing")

	if ptr.Pointer() != nil {
		return int(C.QProxyStyle_LayoutSpacing(ptr.Pointer(), C.int(control1), C.int(control2), C.int(orientation), PointerFromQStyleOption(option), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QProxyStyle) PixelMetric(metric QStyle__PixelMetric, option QStyleOption_ITF, widget QWidget_ITF) int {
	defer qt.Recovering("QProxyStyle::pixelMetric")

	if ptr.Pointer() != nil {
		return int(C.QProxyStyle_PixelMetric(ptr.Pointer(), C.int(metric), PointerFromQStyleOption(option), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QProxyStyle) ConnectPolish(f func(widget *QWidget)) {
	defer qt.Recovering("connect QProxyStyle::polish")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "polish", f)
	}
}

func (ptr *QProxyStyle) DisconnectPolish() {
	defer qt.Recovering("disconnect QProxyStyle::polish")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "polish")
	}
}

//export callbackQProxyStylePolish
func callbackQProxyStylePolish(ptrName *C.char, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QProxyStyle::polish")

	if signal := qt.GetSignal(C.GoString(ptrName), "polish"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QProxyStyle) SetBaseStyle(style QStyle_ITF) {
	defer qt.Recovering("QProxyStyle::setBaseStyle")

	if ptr.Pointer() != nil {
		C.QProxyStyle_SetBaseStyle(ptr.Pointer(), PointerFromQStyle(style))
	}
}

func (ptr *QProxyStyle) SizeFromContents(ty QStyle__ContentsType, option QStyleOption_ITF, size core.QSize_ITF, widget QWidget_ITF) *core.QSize {
	defer qt.Recovering("QProxyStyle::sizeFromContents")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QProxyStyle_SizeFromContents(ptr.Pointer(), C.int(ty), PointerFromQStyleOption(option), core.PointerFromQSize(size), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QProxyStyle) StandardIcon(standardIcon QStyle__StandardPixmap, option QStyleOption_ITF, widget QWidget_ITF) *gui.QIcon {
	defer qt.Recovering("QProxyStyle::standardIcon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QProxyStyle_StandardIcon(ptr.Pointer(), C.int(standardIcon), PointerFromQStyleOption(option), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QProxyStyle) StyleHint(hint QStyle__StyleHint, option QStyleOption_ITF, widget QWidget_ITF, returnData QStyleHintReturn_ITF) int {
	defer qt.Recovering("QProxyStyle::styleHint")

	if ptr.Pointer() != nil {
		return int(C.QProxyStyle_StyleHint(ptr.Pointer(), C.int(hint), PointerFromQStyleOption(option), PointerFromQWidget(widget), PointerFromQStyleHintReturn(returnData)))
	}
	return 0
}

func (ptr *QProxyStyle) SubControlRect(cc QStyle__ComplexControl, option QStyleOptionComplex_ITF, sc QStyle__SubControl, widget QWidget_ITF) *core.QRect {
	defer qt.Recovering("QProxyStyle::subControlRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QProxyStyle_SubControlRect(ptr.Pointer(), C.int(cc), PointerFromQStyleOptionComplex(option), C.int(sc), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QProxyStyle) SubElementRect(element QStyle__SubElement, option QStyleOption_ITF, widget QWidget_ITF) *core.QRect {
	defer qt.Recovering("QProxyStyle::subElementRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QProxyStyle_SubElementRect(ptr.Pointer(), C.int(element), PointerFromQStyleOption(option), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QProxyStyle) ConnectUnpolish(f func(widget *QWidget)) {
	defer qt.Recovering("connect QProxyStyle::unpolish")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "unpolish", f)
	}
}

func (ptr *QProxyStyle) DisconnectUnpolish() {
	defer qt.Recovering("disconnect QProxyStyle::unpolish")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "unpolish")
	}
}

//export callbackQProxyStyleUnpolish
func callbackQProxyStyleUnpolish(ptrName *C.char, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QProxyStyle::unpolish")

	if signal := qt.GetSignal(C.GoString(ptrName), "unpolish"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QProxyStyle) DestroyQProxyStyle() {
	defer qt.Recovering("QProxyStyle::~QProxyStyle")

	if ptr.Pointer() != nil {
		C.QProxyStyle_DestroyQProxyStyle(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}