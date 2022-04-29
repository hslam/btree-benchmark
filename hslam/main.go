package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/hslam/btree"
)

var total int
var degree int
var random bool

func init() {
	flag.IntVar(&total, "total", 10000000, "-total=1000000")
	flag.IntVar(&degree, "degree", 100, "-degree=256")
	flag.BoolVar(&random, "rand", false, "-rand=false")
	flag.Parse()
}

func main() {
	if total < 1 || degree < 2 {
		return
	}
	es := make([]int, total)
	for i := 0; i < total; i++ {
		es[i] = i
	}
	shuffle(es)
	loop := 3
	for i := 0; i < loop; i++ {
		Insert(es)
	}
	for i := 0; i < loop; i++ {
		Search(es)
	}
	for i := 0; i < loop; i++ {
		Delete(es)
	}
	for i := 0; i < loop; i++ {
		Iterate(es)
	}
}

func shuffle(es []int) {
	if random {
		rand.Shuffle(total, func(i, j int) {
			es[i], es[j] = es[j], es[i]
		})
	}
}

func Insert(es []int) {
	tree := btree.New(degree)
	start := time.Now()
	for i := 0; i < total; i++ {
		tree.Insert(btree.Int(es[i]))
	}
	d := time.Now().Sub(start)
	fmt.Printf("Insert:\ttotal %d,\tdegree %d,\ttime %v,\t%d ns/op,\t%d items/s\n", total, degree, d, int(d.Nanoseconds())/total, int(float64(total)*1e9/float64(d.Nanoseconds())))
}

func Search(es []int) {
	tree := btree.New(degree)
	for i := 0; i < total; i++ {
		tree.Insert(btree.Int(es[i]))
	}
	shuffle(es)
	start := time.Now()
	for i := 0; i < total; i++ {
		tree.Search(btree.Int(es[i]))
	}
	d := time.Now().Sub(start)
	fmt.Printf("Search:\ttotal %d,\tdegree %d,\ttime %v,\t%d ns/op,\t%d items/s\n", total, degree, d, int(d.Nanoseconds())/total, int(float64(total)*1e9/float64(d.Nanoseconds())))
}

func Delete(es []int) {
	tree := btree.New(degree)
	for i := 0; i < total; i++ {
		tree.Insert(btree.Int(es[i]))
	}
	shuffle(es)
	start := time.Now()
	for i := 0; i < total; i++ {
		tree.Delete(btree.Int(es[i]))
	}
	d := time.Now().Sub(start)
	fmt.Printf("Delete:\ttotal %d,\tdegree %d,\ttime %v,\t%d ns/op,\t%d items/s\n", total, degree, d, int(d.Nanoseconds())/total, int(float64(total)*1e9/float64(d.Nanoseconds())))
}

func Iterate(es []int) {
	tree := btree.New(degree)
	for i := 0; i < total; i++ {
		tree.Insert(btree.Int(es[i]))
	}
	start := time.Now()
	iter := tree.Min().MinIterator()
	for iter != nil {
		iter.Item()
		iter = iter.Next()
	}
	d := time.Now().Sub(start)
	fmt.Printf("Iterate:\ttotal %d,\tdegree %d,\ttime %v,\t%d ns/op,\t%d items/s\n", total, degree, d, int(d.Nanoseconds())/total, int(float64(total)*1e9/float64(d.Nanoseconds())))
}
