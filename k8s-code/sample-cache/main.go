//__author__ = "YaoYao"
//Date: 2020/9/10
package main

import (
	"k8s.io/client-go/tools/cache"
)

func main() {
	indexers := cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}
	indexer := cache.NewIndexer(cache.DeletionHandlingMetaNamespaceKeyFunc, indexers)
	indexer.Add("xx")
}
