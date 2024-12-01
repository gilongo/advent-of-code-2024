const fs = require("fs");
const path = require("path");

const input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8");

let firstList = [];
let secondList = [];

// split in two lists
input
  .split("\n")
  .map((x) => x.replace("\r", ""))
  .map((x) => {
    newString = x.split("   ");
    try {
      firstList.push(parseInt(newString[0].trim()));
      secondList.push(parseInt(newString[1].trim()));
    } catch (error) {
      console.log(error);
    }
  });

// sort
firstList.sort();
secondList.sort();

// get distance between points of two lists
let listDiff = 0;
for (let i = 0; i < firstList.length; i++) {
  firstMin = firstList[i];
  secondMin = secondList[i];

  listDiff += Math.abs(firstMin - secondMin);
}

console.log("Distance: ", listDiff);

alreadyUsed = {
  0: 0,
};
similarityScore = 0;
firstList.map((x) => {
  if (alreadyUsed[x]) {
    similarityScore += alreadyUsed[x] * x;
    return;
  }

  encountered = 0;
  secondList.map((y) => {
    if (x === y) encountered++;
  });

  alreadyUsed[x] = encountered;
  similarityScore += encountered * x;
});

console.log("similarityScore: ", similarityScore);
