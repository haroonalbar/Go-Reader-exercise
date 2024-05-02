package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (m MyReader) Read(b []byte) (i int , e error){
  for i:= range b{
    b[i]= 'A'
  }
  return len(b), nil

}

func main() {
    reader.Validate(MyReader{})
}

// this was a wired exercise cause the question wasn't clear I had to look it up on stack overflow.
// basically we make a method to set each buffer byte value to "A" and return it's bypte lenght and error
