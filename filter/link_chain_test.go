package filter

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/ytbiu/tool/filter/mock"
)

func TestAppend(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock data here
	mockFilter := mock.NewMockFilter(ctrl)

	LinkFilter := NewLinkFilter(mockFilter)

	linkChain := &LinkChain{}
	linkChain.Append(LinkFilter)

	a.Equal(1, linkChain.size)
	a.Equal(LinkFilter, linkChain.first)
}

func TestRemove(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	{
		// 删除根节点
		name0 := "test_remove_0"
		mockFilter := mock.NewMockFilter(ctrl)
		mockFilter.EXPECT().Name().AnyTimes().Return(name0)

		linkFilter := &LinkFilter{Filter: mockFilter}

		linkChain := &LinkChain{size: 1, first: linkFilter}
		linkChain.Remove(linkFilter)

		a.Equal(0, linkChain.size)
		a.Nil(linkChain.first)
	}
	{

		// 删除第二个节点
		name0 := "test_remove_0"
		mockFilter := mock.NewMockFilter(ctrl)
		mockFilter.EXPECT().Name().AnyTimes().Return(name0)

		name1 := "test_remove_1"
		mockFilter1 := mock.NewMockFilter(ctrl)
		mockFilter1.EXPECT().Name().AnyTimes().Return(name1)

		linkFilter := &LinkFilter{Filter: mockFilter}
		linkFilter1 := &LinkFilter{Filter: mockFilter1}
		linkFilter1.SetNext(linkFilter)

		linkChain1 := &LinkChain{size: 2, first: linkFilter1}
		linkChain1.Remove(mockFilter)

		a.Equal(1, linkChain1.size)
		a.Equal(linkFilter1, linkChain1.first)
		a.Nil(linkChain1.first.Next())
	}

}

func TestCheck(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	{
		// 一个filter 为true
		mockFilter := mock.NewMockFilter(ctrl)
		a.Equal(nil, nil)
		mockFilter.EXPECT().Check(nil).Return(true)

		linkFilter := NewLinkFilter(mockFilter)

		linkChain := &LinkChain{}
		linkChain.Append(linkFilter)
		res := linkChain.Check(nil)
		a.Equal(true, res)
	}

	{
		// 一个filter 为false
		mockFilter := mock.NewMockFilter(ctrl)
		a.Equal(nil, nil)
		mockFilter.EXPECT().Check(nil).Return(false)

		linkFilter := NewLinkFilter(mockFilter)

		linkChain := &LinkChain{}
		linkChain.Append(linkFilter)
		res := linkChain.Check(nil)
		a.Equal(false, res)
	}

	{
		// 2个filter 为false
		mockFilter := mock.NewMockFilter(ctrl)
		a.Equal(nil, nil)
		mockFilter.EXPECT().Check(nil).Return(true)

		mockFilter1 := mock.NewMockFilter(ctrl)
		a.Equal(nil, nil)
		mockFilter1.EXPECT().Check(nil).Return(false)

		linkFilter := NewLinkFilter(mockFilter)
		linkFilter1 := NewLinkFilter(mockFilter1)

		linkChain := &LinkChain{}
		linkChain.Append(linkFilter, linkFilter1)

		res := linkChain.Check(nil)
		a.Equal(false, res)
	}

}
