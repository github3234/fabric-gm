// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/tw-bc-group/fabric-gm/core/ledger"
	"github.com/tw-bc-group/fabric-gm/protos/common"
	"github.com/tw-bc-group/fabric-gm/protos/ledger/rwset/kvrwset"
)

type DeployedChaincodeInfoProvider struct {
	NamespacesStub        func() []string
	namespacesMutex       sync.RWMutex
	namespacesArgsForCall []struct{}
	namespacesReturns     struct {
		result1 []string
	}
	namespacesReturnsOnCall map[int]struct {
		result1 []string
	}
	UpdatedChaincodesStub        func(stateUpdates map[string][]*kvrwset.KVWrite) ([]*ledger.ChaincodeLifecycleInfo, error)
	updatedChaincodesMutex       sync.RWMutex
	updatedChaincodesArgsForCall []struct {
		stateUpdates map[string][]*kvrwset.KVWrite
	}
	updatedChaincodesReturns struct {
		result1 []*ledger.ChaincodeLifecycleInfo
		result2 error
	}
	updatedChaincodesReturnsOnCall map[int]struct {
		result1 []*ledger.ChaincodeLifecycleInfo
		result2 error
	}
	ChaincodeInfoStub        func(chaincodeName string, qe ledger.SimpleQueryExecutor) (*ledger.DeployedChaincodeInfo, error)
	chaincodeInfoMutex       sync.RWMutex
	chaincodeInfoArgsForCall []struct {
		chaincodeName string
		qe            ledger.SimpleQueryExecutor
	}
	chaincodeInfoReturns struct {
		result1 *ledger.DeployedChaincodeInfo
		result2 error
	}
	chaincodeInfoReturnsOnCall map[int]struct {
		result1 *ledger.DeployedChaincodeInfo
		result2 error
	}
	CollectionInfoStub        func(chaincodeName, collectionName string, qe ledger.SimpleQueryExecutor) (*common.StaticCollectionConfig, error)
	collectionInfoMutex       sync.RWMutex
	collectionInfoArgsForCall []struct {
		chaincodeName  string
		collectionName string
		qe             ledger.SimpleQueryExecutor
	}
	collectionInfoReturns struct {
		result1 *common.StaticCollectionConfig
		result2 error
	}
	collectionInfoReturnsOnCall map[int]struct {
		result1 *common.StaticCollectionConfig
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DeployedChaincodeInfoProvider) Namespaces() []string {
	fake.namespacesMutex.Lock()
	ret, specificReturn := fake.namespacesReturnsOnCall[len(fake.namespacesArgsForCall)]
	fake.namespacesArgsForCall = append(fake.namespacesArgsForCall, struct{}{})
	fake.recordInvocation("Namespaces", []interface{}{})
	fake.namespacesMutex.Unlock()
	if fake.NamespacesStub != nil {
		return fake.NamespacesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.namespacesReturns.result1
}

func (fake *DeployedChaincodeInfoProvider) NamespacesCallCount() int {
	fake.namespacesMutex.RLock()
	defer fake.namespacesMutex.RUnlock()
	return len(fake.namespacesArgsForCall)
}

func (fake *DeployedChaincodeInfoProvider) NamespacesReturns(result1 []string) {
	fake.NamespacesStub = nil
	fake.namespacesReturns = struct {
		result1 []string
	}{result1}
}

func (fake *DeployedChaincodeInfoProvider) NamespacesReturnsOnCall(i int, result1 []string) {
	fake.NamespacesStub = nil
	if fake.namespacesReturnsOnCall == nil {
		fake.namespacesReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.namespacesReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *DeployedChaincodeInfoProvider) UpdatedChaincodes(stateUpdates map[string][]*kvrwset.KVWrite) ([]*ledger.ChaincodeLifecycleInfo, error) {
	fake.updatedChaincodesMutex.Lock()
	ret, specificReturn := fake.updatedChaincodesReturnsOnCall[len(fake.updatedChaincodesArgsForCall)]
	fake.updatedChaincodesArgsForCall = append(fake.updatedChaincodesArgsForCall, struct {
		stateUpdates map[string][]*kvrwset.KVWrite
	}{stateUpdates})
	fake.recordInvocation("UpdatedChaincodes", []interface{}{stateUpdates})
	fake.updatedChaincodesMutex.Unlock()
	if fake.UpdatedChaincodesStub != nil {
		return fake.UpdatedChaincodesStub(stateUpdates)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updatedChaincodesReturns.result1, fake.updatedChaincodesReturns.result2
}

func (fake *DeployedChaincodeInfoProvider) UpdatedChaincodesCallCount() int {
	fake.updatedChaincodesMutex.RLock()
	defer fake.updatedChaincodesMutex.RUnlock()
	return len(fake.updatedChaincodesArgsForCall)
}

func (fake *DeployedChaincodeInfoProvider) UpdatedChaincodesArgsForCall(i int) map[string][]*kvrwset.KVWrite {
	fake.updatedChaincodesMutex.RLock()
	defer fake.updatedChaincodesMutex.RUnlock()
	return fake.updatedChaincodesArgsForCall[i].stateUpdates
}

func (fake *DeployedChaincodeInfoProvider) UpdatedChaincodesReturns(result1 []*ledger.ChaincodeLifecycleInfo, result2 error) {
	fake.UpdatedChaincodesStub = nil
	fake.updatedChaincodesReturns = struct {
		result1 []*ledger.ChaincodeLifecycleInfo
		result2 error
	}{result1, result2}
}

func (fake *DeployedChaincodeInfoProvider) UpdatedChaincodesReturnsOnCall(i int, result1 []*ledger.ChaincodeLifecycleInfo, result2 error) {
	fake.UpdatedChaincodesStub = nil
	if fake.updatedChaincodesReturnsOnCall == nil {
		fake.updatedChaincodesReturnsOnCall = make(map[int]struct {
			result1 []*ledger.ChaincodeLifecycleInfo
			result2 error
		})
	}
	fake.updatedChaincodesReturnsOnCall[i] = struct {
		result1 []*ledger.ChaincodeLifecycleInfo
		result2 error
	}{result1, result2}
}

func (fake *DeployedChaincodeInfoProvider) ChaincodeInfo(chaincodeName string, qe ledger.SimpleQueryExecutor) (*ledger.DeployedChaincodeInfo, error) {
	fake.chaincodeInfoMutex.Lock()
	ret, specificReturn := fake.chaincodeInfoReturnsOnCall[len(fake.chaincodeInfoArgsForCall)]
	fake.chaincodeInfoArgsForCall = append(fake.chaincodeInfoArgsForCall, struct {
		chaincodeName string
		qe            ledger.SimpleQueryExecutor
	}{chaincodeName, qe})
	fake.recordInvocation("ChaincodeInfo", []interface{}{chaincodeName, qe})
	fake.chaincodeInfoMutex.Unlock()
	if fake.ChaincodeInfoStub != nil {
		return fake.ChaincodeInfoStub(chaincodeName, qe)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.chaincodeInfoReturns.result1, fake.chaincodeInfoReturns.result2
}

func (fake *DeployedChaincodeInfoProvider) ChaincodeInfoCallCount() int {
	fake.chaincodeInfoMutex.RLock()
	defer fake.chaincodeInfoMutex.RUnlock()
	return len(fake.chaincodeInfoArgsForCall)
}

func (fake *DeployedChaincodeInfoProvider) ChaincodeInfoArgsForCall(i int) (string, ledger.SimpleQueryExecutor) {
	fake.chaincodeInfoMutex.RLock()
	defer fake.chaincodeInfoMutex.RUnlock()
	return fake.chaincodeInfoArgsForCall[i].chaincodeName, fake.chaincodeInfoArgsForCall[i].qe
}

func (fake *DeployedChaincodeInfoProvider) ChaincodeInfoReturns(result1 *ledger.DeployedChaincodeInfo, result2 error) {
	fake.ChaincodeInfoStub = nil
	fake.chaincodeInfoReturns = struct {
		result1 *ledger.DeployedChaincodeInfo
		result2 error
	}{result1, result2}
}

func (fake *DeployedChaincodeInfoProvider) ChaincodeInfoReturnsOnCall(i int, result1 *ledger.DeployedChaincodeInfo, result2 error) {
	fake.ChaincodeInfoStub = nil
	if fake.chaincodeInfoReturnsOnCall == nil {
		fake.chaincodeInfoReturnsOnCall = make(map[int]struct {
			result1 *ledger.DeployedChaincodeInfo
			result2 error
		})
	}
	fake.chaincodeInfoReturnsOnCall[i] = struct {
		result1 *ledger.DeployedChaincodeInfo
		result2 error
	}{result1, result2}
}

func (fake *DeployedChaincodeInfoProvider) CollectionInfo(chaincodeName string, collectionName string, qe ledger.SimpleQueryExecutor) (*common.StaticCollectionConfig, error) {
	fake.collectionInfoMutex.Lock()
	ret, specificReturn := fake.collectionInfoReturnsOnCall[len(fake.collectionInfoArgsForCall)]
	fake.collectionInfoArgsForCall = append(fake.collectionInfoArgsForCall, struct {
		chaincodeName  string
		collectionName string
		qe             ledger.SimpleQueryExecutor
	}{chaincodeName, collectionName, qe})
	fake.recordInvocation("CollectionInfo", []interface{}{chaincodeName, collectionName, qe})
	fake.collectionInfoMutex.Unlock()
	if fake.CollectionInfoStub != nil {
		return fake.CollectionInfoStub(chaincodeName, collectionName, qe)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.collectionInfoReturns.result1, fake.collectionInfoReturns.result2
}

func (fake *DeployedChaincodeInfoProvider) CollectionInfoCallCount() int {
	fake.collectionInfoMutex.RLock()
	defer fake.collectionInfoMutex.RUnlock()
	return len(fake.collectionInfoArgsForCall)
}

func (fake *DeployedChaincodeInfoProvider) CollectionInfoArgsForCall(i int) (string, string, ledger.SimpleQueryExecutor) {
	fake.collectionInfoMutex.RLock()
	defer fake.collectionInfoMutex.RUnlock()
	return fake.collectionInfoArgsForCall[i].chaincodeName, fake.collectionInfoArgsForCall[i].collectionName, fake.collectionInfoArgsForCall[i].qe
}

func (fake *DeployedChaincodeInfoProvider) CollectionInfoReturns(result1 *common.StaticCollectionConfig, result2 error) {
	fake.CollectionInfoStub = nil
	fake.collectionInfoReturns = struct {
		result1 *common.StaticCollectionConfig
		result2 error
	}{result1, result2}
}

func (fake *DeployedChaincodeInfoProvider) CollectionInfoReturnsOnCall(i int, result1 *common.StaticCollectionConfig, result2 error) {
	fake.CollectionInfoStub = nil
	if fake.collectionInfoReturnsOnCall == nil {
		fake.collectionInfoReturnsOnCall = make(map[int]struct {
			result1 *common.StaticCollectionConfig
			result2 error
		})
	}
	fake.collectionInfoReturnsOnCall[i] = struct {
		result1 *common.StaticCollectionConfig
		result2 error
	}{result1, result2}
}

func (fake *DeployedChaincodeInfoProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.namespacesMutex.RLock()
	defer fake.namespacesMutex.RUnlock()
	fake.updatedChaincodesMutex.RLock()
	defer fake.updatedChaincodesMutex.RUnlock()
	fake.chaincodeInfoMutex.RLock()
	defer fake.chaincodeInfoMutex.RUnlock()
	fake.collectionInfoMutex.RLock()
	defer fake.collectionInfoMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DeployedChaincodeInfoProvider) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ ledger.DeployedChaincodeInfoProvider = new(DeployedChaincodeInfoProvider)
