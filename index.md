# Using a neural network in Go

#go #neuralnetwork #machinelearning

## NNHelper is a Go package for creating and using a neural network

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/unv7ad995nng6wsgmv3l.png)

## Summary

This article describes the work of the [nnhelper](https://github.com/kirill-scherba/nnhelper) package, designed to create and use neural networks in Go programs.

If you are already familiar with machine learning and use it in your work, then this article and the examples described in it may seem too simple for you. If you are at the beginning of your journey and want to get acquainted with this topic or would like to learn how to use the neural network in your Go programs, then you have come to the right place.

The [nnhelper](https://github.com/kirill-scherba/nnhelper) Go package is designed to quickly create a neural network and use it in the applications written in the Go language. To use [nnhelper](https://github.com/kirill-scherba/nnhelper), you don't need anything else other than Go. The [nnhelper](https://github.com/kirill-scherba/nnhelper) package is an add-on to the [gonn](https://github.com/fxsjy/gonn) package. And this is the only one external dependency.

## Neural network (neural matrix)

Let me try to explain in my own words what a neural network or neural matrix is, in order not to copy here complex and, perhaps, incomprehensible explanations from the initial reading.

A neural matrix is an array of floating point values that allows you to give the desired output values depending on the input values. The values known to us are fed to the input of the neural matrix, and the results we expect are given at the output.

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/j6ngzyv1kzl4aorplm70.png)

Let's use the neural matrix to solve the problem "Two Minuses Make a Plus." To do this, we will make two tables: one with the input data, the other with the results.

We have two input parameters and two output parameters. The input is two numbers with a plus or minus sign. At the output we get two numbers with values from 0 to 1. If the first value is close to 1, then the result is Plus, if the second value is close to 1, then the result is Minus. That is, we give real values as input, and we get an array of results as output. Than we select the largest value from the array and consider it the answer.

```
Input data:              Results(plus, minus):
 1, 1  – plus * plus     1, 0 – plus
 1,-1  – plus * minus    0, 1 – minus
-1, 1  – minus* plus     0, 1 – minus
-1,-1  – minus* minus    1, 0 – plus
```

From the prepared tables with input and output data, a neural matrix is created by its training: The input data is fed to the inputs of the matrix. The matrix produces an output that is compared with the result from our results table. Based on these answers, the coefficients in the matrix change. The input data is looped through hundreds, thousands and even hundreds of thousands of iterations until the matrix answers are accurate enough.

## Let's move on to programming

Enough words and theory. Let's see how it looks in practice.

Let's create a folder for the project. And create three files in it:

```
main.go 
sam03_inp.csv 
sam03_tar.csv
```

Let's place our input and output data in the *.csv files:

sam03_inp.csv

```
1,1
1,-1
-1,1
-1,-1
```

sam03_tar.csv

```
1,0
0,1
0,1
1,0
```

In the main.go file, create the main function and write the following code:

```go
// Constants with filenames of our matrix and data
const (
    SAM03_NN  = "sam03.nn"
    SAM03_INP = "sam03_inp.csv"
    SAM03_TAR = "sam03_tar.csv"
)

// Human answers string array
humanAnswers := []string{"Plus", "Minus"}

// Create NN if it file does not exists
if _, err := os.Stat(SAM03_NN); errors.Is(err, os.ErrNotExist) {
        log.Println("Create", SAM03_NN, "neural network")
        nnhelper.Create(2, 4, 2, false, SAM03_INP, SAM03_TAR, SAM03_NN, true)
}

// Load neural matrix from file
nn := nnhelper.Load(SAM03_NN)

// Using / testing our neural network
const (
        PLUS  = 1.0
        MINUS = -1.0
    )

    // Intput array for testing
    in := [][]float64{
        {PLUS, PLUS},   // Plus * Plus = Plus
        {PLUS, MINUS},  // Plus * Minus = Minus
        {MINUS, PLUS},  // Minus * Plus = Minus
        {MINUS, MINUS}, // Minus * Minus = Plus
        {3000, -0.001}, // Minus * Plus = Minus
    }
    for i := range in {
        out := nn.Answer(in[i]...)
        answer, _ := nn.AnswerToHuman(out, humanAnswers)
        fmt.Println(in[i], answer, out)
    }
}
```

The full text of this example and the data files are located in [examples/sam03](https://github.com/kirill-scherba/nnhelper/blob/main/examples/sam03)

Let's run the example:

```
go run .
```

And we get the results:

```
[1 1] Plus [0.9944239772210877 0.005449692189449571]
[1 -1] Minus [0.006860785779850435 0.9935960167863507]
[-1 1] Minus [0.005651009980489101 0.994384581174021]
[-1 -1] Plus [0.9944591181959666 0.005221796400203198]
[3000 -0.001] Minus [0.005445102841471242 0.9960123783099599]
```

In the results we see (see the first line):

- our original data: [1,1]
- result translated into understandable form: Plus
- result obtained from matrix outputs: [0.9944239772210877 0.005449692189449571]

Pay attention to the last line of the results. Our matrix was able to give the correct answer to the "unknown" inputs. This is the whole taste of neural networks. The matrix gives answers not only to the input data that it was trained on, but also to other input parameters unknown to it.

To see the training process of the neural network, you need to delete the sam03.nn file and run the example again.

Well, I guess I'll wrap it up here. For a long time I dreamed of writing a clear and simple explanation of how neural networks can be used in programming. I hope I succeeded.

There are two more examples in the package:

- _sam02_: Input the time in 24-hour format and get the answer: Morning, Evening, Day or Night;
- _sam01_: An example of a matrix for getting the reaction of a game bot. The input is the amount of health, the presence of weapons, the number of enemies, and at the output we get the answer to the question “what to do”: attack, sneak, run away or do nothing.

The package is hosted on Github:  
https://github.com/kirill-scherba/nnhelper

Best regards,  
Kirill Scherba
