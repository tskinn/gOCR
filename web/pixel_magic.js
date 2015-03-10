
var letters = document.querySelectorAll(".weights");
var colorSpectrum = ["#0515DC", "#1A45BF", "#378399", "#4BAD7E", "#5BCF6A", "#81A050", "#958643", "#BF5028", "#E91A0D", "#FB0301"];

console.log(letters.length);


function changeColor (letter, row, col, color) {
    var letters = document.querySelectorAll(".weights")[letter];
    var rows = letters.querySelectorAll(".weight_row")[row];
    var pixel = rows.querySelectorAll(".weight_pixel")[col];
    pixel.style.backgroundColor = color;
}

// given a double value convert to color value
function weightToColorValue(value) {
    //value = value + 0.5;
    value = value * 10;
    return colorSpectrum[(value | 0)]; // or parseInt(value, 10)
}

function  updateWeightMapColor (weights) {
    for (letter = 0; letter < weights.length; letter++) {
        for (row = 0; row < weights[letter].length; row++) {
            for (col = 0; col < weights[letter][row].length; col++) {
                changeColor(letter, row, col, weightToColorValue(weights[letter][row][col]))
            }
        }
    }
}

var socket = new WebSocket("ws://localhost:3000/ws");

socket.onopen = function (event) {
    console.log("Connection succesfullll!");
};

socket.onmessage = function (event) {
    updateWeightMapColor(JSON.parse(event.data));
    console.log(event.data);
};
