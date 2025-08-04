package registry

// To perform the test, start a local registry server at port 5000
//  docker run -it --rm -p 5000:5000 /v host/dir:container/dir registry:2.8
/*
import (
	"testing"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func Test_CacheImage_with_server(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		//src_image := "mcr.microsoft.com/powershell:7.5-ubuntu-noble"
		//src_image := "alpine:3.19"
		src_image := "alpine:3"
		Endpoint = "localhost:5000"
		sourceRef, err := name.ParseReference(src_image)
		desc, err := remote.Get(sourceRef)
		onUpdated := func(update v1.Update) {
			t.Log("Progress Update -", update)
		}

		onComplete := func(totalSize int64) {
			t.Log("Total Size -", totalSize)
		}
		t.Log("Initial cache")
		err = CacheImage(src_image, desc, []string{"amd64"}, onUpdated, onComplete)
		if err != nil {
			t.Log("Initial cache failed.", err)
			panic("Failed to perform initial cache image")
		}
	})

}
*/
