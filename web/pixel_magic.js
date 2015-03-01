
var letters = document.querySelectorAll(".letter");
var colorSpectrum = ["#0515DC", "#1A45BF", "#378399", "#4BAD7E", "#5BCF6A", "#81A050", "#958643", "#BF5028", "#E91A0D", "#FB0301"];

console.log(letters.length);


function changeColor (letter, row, col, color) {
    var letters = document.querySelectorAll(".letter")[letter];
    var rows = letters.querySelectorAll(".row")[row];
    var pixel = rows.querySelectorAll(".pixel")[col];
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

var socket = new WebSocket("ws://localhost:3000/ws", "protocol");

socket.onopen = function (event) {
    console.log("Connection succesfullll!");
};


socket.onmessage = function (event) {
    updateWeightMapColor(JSON.parse(event.data));
    console.log(event.data);
};


// function change(color) {
//    for (i = 0;i < 10; i++) {
//       for (j = 0; j < 10; j++) {
//          changeColor("a", i, j, color);
//          changeColor("b", i, j, color);
//          changeColor("c", i, j, color);
//          changeColor("d", i, j, color);
//          changeColor("e", i, j, color);
//          changeColor("f", i, j, color);
//          changeColor("g", i, j, color);
//          changeColor("h", i, j, color);
//          changeColor("i", i, j, color);
//          changeColor("j", i, j, color);
//          changeColor("k", i, j, color);
//          changeColor("l", i, j, color);
//          changeColor("m", i, j, color);
//          changeColor("n", i, j, color);
//          changeColor("o", i, j, color);
//          changeColor("p", i, j, color);
//          changeColor("q", i, j, color);
//          changeColor("r", i, j, color);
//          changeColor("s", i, j, color);
//          changeColor("t", i, j, color);
//          changeColor("u", i, j, color);
//          changeColor("v", i, j, color);
//          changeColor("w", i, j, color);
//          changeColor("x", i, j, color);
//          changeColor("y", i, j, color);
//          changeColor("z", i, j, color);
//       }
//    }
// }

// function changeColorssss() {
//    console.log(document.getElementById("myColor").value);
//    changeColors(document.getElementById("myColor").value);
// }

// function changeColors (color) {
//    var letters = document.querySelectorAll(".letter");
//    for (i = 0; i < letters.length; i++) {
//       var rows = letters[i].querySelectorAll(".row");
//       for (j = 0; j < rows.length; j++) {
//          var pixels = rows[j].querySelectorAll(".pixel");
//          for (k = 0; k < pixels.length; k++) {
//             pixels[k].style.backgroundColor = color;
//          }
//       }
//    }
// }
