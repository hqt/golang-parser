package example

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRunGolangParser(t *testing.T) {
	rawListener := RunGolangParser()
	listener, ok := rawListener.(*GoParserListenerImpl)
	require.True(t, ok)
	require.Equal(t, 2, len(listener.methods))
	for i := 0; i < len(listener.methods); i++ {
		method := listener.methods[i]
		require.Equal(t, len(method.Parameters), len(method.Types))
		fmt.Println(method)
	}
}
