const fs = require("fs");
const path = require("path");

const baseDir = __dirname;
const outputFile = path.join(baseDir, "output.txt");

function collectFiles(dirPath) {
  let filesList = [];
  const items = fs.readdirSync(dirPath, { withFileTypes: true });

  for (const item of items) {
    const fullPath = path.join(dirPath, item.name);

    // Skip node_modules folders
    if (item.isDirectory() && item.name === "node_modules") {
      continue;
    }

    if (item.isDirectory()) {
      filesList = filesList.concat(collectFiles(fullPath));
    } else if (
      item.isFile() &&
      item.name !== "README.md" &&
      item.name !== "package-lock.json"
    ) {
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

  const folders = fs.readdirSync(baseDir, { withFileTypes: true })
    .filter(entry => entry.isDirectory())
    .map(entry => entry.name);

  for (const folder of folders) {
    const fullPath = path.join(baseDir, folder);
    const files = collectFiles(fullPath);
    for (const file of files) {
      writeFileOutput(file, outputStream);
    }
  }

  outputStream.end(() => {
    console.log("✅ output.txt generated successfully.");
  });
}

main();
