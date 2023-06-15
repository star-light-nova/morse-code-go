package reversemap_test

import (
    "testing"
    "github.com/StarLightNova/Morse-Code-Go/pkg/reversemap"
)

func TestReverseMap(t *testing.T) {
    originalMap := map[string]string{"a": "b"}
    resultMap := map[string]string{"b": "a"}

    testMap := reversemap.Reverse(originalMap)

    for _, tVal := range testMap {
        if tVal != resultMap["b"] {
            t.Fatal("The Reverse Map didn't reversed the map.")
        }
    }
}
