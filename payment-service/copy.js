const fs = require("fs");
const path = require("path");

const baseDir = __dirname;
const foldersToScan = ["cmd", "internal", "tests"];
const outputFile = path.join(baseDir, "output.txt");

function collectFiles(dirPath) {
  let filesList = [];

  const items = fs.readdirSync(dirPath, { withFileTypes: true });
  for (const item of items) {
    const fullPath = path.join(dirPath, item.name);
    if (item.isDirectory()) {
      filesList = filesList.concat(collectFiles(fullPath));
    } else if (item.isFile()) {
      filesList.push(fullPath);
    }
  }

  return filesList;
}

function writeFileOutput(filePath, stream) {
  const relativePath = path.relative(baseDir, filePath);
  const content = fs.readFileSync(filePath, "utf-8");
  stream.write(`=== ${relativePath} ===\n`);
  stream.write(content + "\n\n");
}

function main() {
  const outputStream = fs.createWriteStream(outputFile, { flags: "w" });

  // Process folders
  for (const folder of foldersToScan) {
    const fullPath = path.join(baseDir, folder);
    if (fs.existsSync(fullPath)) {
      const files = collectFiles(fullPath);
      for (const file of files) {
        writeFileOutput(file, outputStream);
      }
    }
  }

  // Add go.mod and go.sum
  for (const fileName of ["go.mod"]) {
    const fullPath = path.join(baseDir, fileName);
    if (fs.existsSync(fullPath)) {
      writeFileOutput(fullPath, outputStream);
    }
  }

  outputStream.end(() => {
    console.log("✅ output.txt generated successfully.");
  });
}

main();
