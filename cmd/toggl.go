// Copyright (c) 2017 Tao Chen <pro711@gmail.com>

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

package cmd

import (
    "os"
    "fmt"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var TogglCommand = &cobra.Command{
    Use: "toggl",
    Short: "A command line Toggl client",
}

func Execute() {
    if err := TogglCommand.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}

func init() {
    cobra.OnInitialize(initConfig)
}

func initConfig() {
    viper.SetConfigName(".toggl") // name of config file (without extension)
    viper.AddConfigPath("$HOME")  // call multiple times to add many search paths

    if err := viper.ReadInConfig(); err != nil {
        fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}
