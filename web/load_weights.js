

function insertDivs(parentclassname, childClassName, numberOfRows) {
    var parents = document.querySelectorAll("." + parentclassname);
    for (var i = 0; i < parents.length; i++) {
        for (var j = 0; j < numberOfRows; j++) {
            var row = document.createElement("div");
            row.className = childClassName;
            parents[i].appendChild(row);
        }
    }
}

insertDivs("weights", "weight_row", 9);
insertDivs("weight_row", "weight_pixel", 9);

insertDivs("char", "char_row", 9);
insertDivs("char_row", "char_pix", 9);

insertDivs("customc", "c_row", 9);
insertDivs("c_row", "c_pix", 9);

console.log("done loading pixels");

