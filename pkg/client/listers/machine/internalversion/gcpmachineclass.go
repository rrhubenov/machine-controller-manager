// This file was automatically generated by lister-gen

package internalversion

import (
	machine "github.com/gardener/node-controller-manager/pkg/apis/machine"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GCPMachineClassLister helps list GCPMachineClasses.
type GCPMachineClassLister interface {
	// List lists all GCPMachineClasses in the indexer.
	List(selector labels.Selector) (ret []*machine.GCPMachineClass, err error)
	// Get retrieves the GCPMachineClass from the index for a given name.
	Get(name string) (*machine.GCPMachineClass, error)
	GCPMachineClassListerExpansion
}

// gCPMachineClassLister implements the GCPMachineClassLister interface.
type gCPMachineClassLister struct {
	indexer cache.Indexer
}

// NewGCPMachineClassLister returns a new GCPMachineClassLister.
func NewGCPMachineClassLister(indexer cache.Indexer) GCPMachineClassLister {
	return &gCPMachineClassLister{indexer: indexer}
}

// List lists all GCPMachineClasses in the indexer.
func (s *gCPMachineClassLister) List(selector labels.Selector) (ret []*machine.GCPMachineClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.GCPMachineClass))
	})
	return ret, err
}

// Get retrieves the GCPMachineClass from the index for a given name.
func (s *gCPMachineClassLister) Get(name string) (*machine.GCPMachineClass, error) {
	key := &machine.GCPMachineClass{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(machine.Resource("gcpmachineclass"), name)
	}
	return obj.(*machine.GCPMachineClass), nil
}
