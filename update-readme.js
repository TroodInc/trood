const fs = require('fs');
const path = require('path');

// Paths to the .gitmodules and repos.json files
const gitmodulesPath = path.join(__dirname, '.gitmodules');
const reposJsonPath = path.join(__dirname, 'repos.json');
const readmePath = path.join(__dirname, 'README.md');

// Function to parse .gitmodules
function parseGitmodules() {
  const gitmodules = fs.readFileSync(gitmodulesPath, 'utf8');
  const submodules = [];
  const regex = /\[submodule "(.*?)"\]\s+path = (.*?)\s+url = (.*?)/g;
  let match;
  while ((match = regex.exec(gitmodules)) !== null) {
    submodules.push({
      name: match[1],
      path: match[2],
      url: match[3]
    });
  }
  return submodules;
}

// Function to load descriptions from repos.json
function loadDescriptions() {
  if (fs.existsSync(reposJsonPath)) {
    const data = fs.readFileSync(reposJsonPath, 'utf8');
    return JSON.parse(data);
  } else {
    console.warn('repos.json not found, proceeding without descriptions.');
    return {};
  }
}

// Function to generate README content
function generateReadme(submodules, descriptions) {
  let content = `# Trood Metarepository\n\nThis repository contains the following submodules:\n\n`;
  content += `| Repository Name | Path | URL | Description |\n| --- | --- | --- | --- |\n`;

  submodules.forEach(submodule => {
    const description = descriptions[submodule.name]?.description || 'No description available.';
    content += `| **${submodule.name}** | \`${submodule.path}\` | [${submodule.url}](${submodule.url}) | ${description} |\n`;
  });

  return content;
}

// Main function
function updateReadme() {
  const submodules = parseGitmodules();
  const descriptions = loadDescriptions();
  const newReadmeContent = generateReadme(submodules, descriptions);

  // Write the new README content to README.md
  fs.writeFileSync(readmePath, newReadmeContent, 'utf8');
}

// Run the update function
updateReadme();
