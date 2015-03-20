
var letters = document.querySelectorAll(".weights");
var colorSpectrum = ["#0515DC", "#1A45BF", "#378399", "#4BAD7E", "#5BCF6A", "#81A050", "#958643", "#BF5028", "#E91A0D", "#FB0301"];

var pixelColors = ["rgb(145, 170, 157)", "rgb(25, 52, 65)"];

console.log(letters.length);


function changeColor (letter, row, col, color) {
    var letters = document.querySelectorAll(".weights")[letter];
    var rows = letters.querySelectorAll(".weight_row")[row];
    var pixel = rows.querySelectorAll(".weight_pixel")[col];
    pixel.style.backgroundColor = color;
}                                                                   //TODO combine these two functions

function changeColorLetter (letter, row, col, color) {
    var letters = document.querySelectorAll(".char")[letter];
    var rows = letters.querySelectorAll(".char_row")[row];
    var pixel = rows.querySelectorAll(".char_pix")[col];
    pixel.style.backgroundColor = color;
}


// given a double value convert to color value
function weightToColorValue(value) {
    //value = value + 0.5;
    value = value * 10;
    return colorSpectrum[(value | 0)]; // or parseInt(value, 10)
}

function updateLetterColor(value) {
    return pixelColors[value];
    // if (value == 1) {
    //     return "rgb(25, 52, 65)"
    // } else {
    //     return "rgb(145, 170, 157)";
    // }
}

function updateLetters(letters, winners) {
    var charnames = document.querySelectorAll(".charname");
    for (letter = 0; letter < letters.length; letter++) {
        for (row = 0; row < letters[letter].length; row++) {
            for (col = 0; col < letters[letter][row].length; col++) {
                changeColorLetter(letter, row, col, updateLetterColor(letters[letter][row][col]))
            }
        }
        
        charnames.item(letter).innerHTML = winners[letter].toUpperCase();
    }
}

function  updateWeightMapColor (weights) {
    for (letter = 0; letter < weights.length; letter++) {
        for (row = 0; row < weights[letter].length; row++) {
            for (col = 0; col < weights[letter][row].length; col++) {
                changeColor(letter, row, col, weightToColorValue(weights[letter][row][col]));
            }
        }
    }
}

function getLetters () {
    ;
}

var message = {
    message: null,
    isInitialized: null,
    neuralNet: null,
    letters: null,
    winners: null,
    numLetters: null,
    learningRate: null,
    totalIterations: null,
    currentIteration: null,
    updateInterval: null,
    neighborEffect: null
};

function printMessage() {
    console.log("Message:      " + message.message);
    console.log("NumLetters:   " + message.numLetters);
    console.log("LearningRate: " + message.learningRate);
    console.log("TotalIters:   " + message.totalIterations);
    console.log("UpdateInter:  " + message.updateInterval);
    console.log("NeighborEfct: " + message.neighborEffect);
}

function train() {
    message.message = "start";
    message.numLetters = parseInt(document.getElementById("numletters").value);
    message.learningRate = parseFloat(document.getElementById("learningrate").value);
    message.totalIterations = parseFloat(document.getElementById("iterations").value);
    message.updateInterval = parseInt(document.getElementById("updateinterval").value);
    //message.updateInterval = parseFloat(document.getElementById("iterations").value);
    message.neighborEffect = parseFloat(document.getElementById("neighboreffect").value);
    socket.send(JSON.stringify(message));
    //printMessage();
}

function continueTraining() {
    message.message = "continue";
    message.updateInterval = parseInt(document.getElementById("updateinterval").value);
    socket.send(JSON.stringify(message));
}

function reset() {
    message.message = "reset";
    socket.send(JSON.stringify(message));
    console.log("reset");
}

var socket = new WebSocket("ws://localhost:3000/ws");

socket.onopen = function (event) {
    console.log("Successfully connected to " + socket.url);
};

socket.onmessage = function (event) {
    message = JSON.parse(event.data);
    console.log("Message: " + message.message);
    if (message.message == "update") {        
        updateWeightMapColor(message.neuralNet);
        message.message = "continue";
        socket.send(JSON.stringify(message));
    } else if (message.message == "done") {
        ;
    } else if (message.message == "init") {
        updateWeightMapColor(message.neuralNet);
    }
    
    updateLetters(message.letters, message.winners);
    //console.log(event.data);
};
