const fs = require("fs");
const path = require("path");

// Paths to the .gitmodules and repos.json files
const gitmodulesPath = path.join(__dirname, ".gitmodules");
const reposJsonPath = path.join(__dirname, "repos.json");
const readmePath = path.join(__dirname, "README.md");

// Function to parse .gitmodules
function parseGitmodules() {
  const gitmodules = fs.readFileSync(gitmodulesPath, "utf8");
  const submodules = [];
  const regex = /\[submodule "(.*?)"\]\s+path = (.*?)\s+url = (.*?)/g;
  let match;
  while ((match = regex.exec(gitmodules)) !== null) {
    submodules.push({
      name: match[1],
      path: match[2],
      url: match[3],
    });
  }
  return submodules;
}

// Function to load descriptions from repos.json
function loadDescriptions() {
  if (fs.existsSync(reposJsonPath)) {
    const data = fs.readFileSync(reposJsonPath, "utf8");
    return JSON.parse(data);
  } else {
    console.warn("repos.json not found, proceeding without descriptions.");
    return {};
  }
}

// Function to generate README content
function generateReadme(submodules, descriptions) {
  let content = `# Trood Metarepository\n\n`;

  // Introductory section with the elevator pitch
  content += `Welcome to the Trood Metarepository! Here, we gather and share tools, libraries, and resources to empower developers worldwide to build faster, better, and more efficiently.\n\n`;
  content += `Trood is an AI-powered technical director that manages the launch and further development of custom software applications. We automate critical processes like market research, product documentation, design, code generation, deployment, and integration—helping businesses and individuals bring their ideas to life with less risk and greater efficiency.\n\n`;
  content += `For the community, Trood provides a platform to collaborate, share, and showcase innovative projects while enabling developers to contribute to something greater. By participating, you can:\n`;
  content += `- Connect with like-minded peers and teams worldwide.\n`;
  content += `- Gain recognition for your work and expertise.\n`;
  content += `- Contribute to open-source projects that are changing how software is built.\n\n`;

  // Encouragement for developers
  content += `Join us in building the future of software development. Whether you're sharing your own libraries, contributing to existing projects, or learning from others, there's a place for you here.\n\n`;

  // Table of submodules
  content += `## Submodules\n\nThis repository contains the following submodules:\n\n`;
  content += `| Repository Name | Path | URL | Description |\n| --- | --- | --- | --- |\n`;

  submodules.forEach((submodule) => {
    const description =
      descriptions[submodule.name].description || "No description available.";
    // const description = 'No description available.';
    content += `| **${submodule.name}** | \`${submodule.path}\` | [${submodule.name}](${submodule.url}) | ${description} |\n`;
  });

  // Call-to-action with links
  content += `\n## Get Involved\n\n`;
  content += `We believe in the power of collaboration. Here's how you can participate:\n`;
  content += `- **Contribute**: Submit pull requests, report issues, and improve the libraries.\n`;
  content += `- **Share**: Publish your own libraries in the Trood Metarepository to reach a global audience.\n`;
  content += `- **Collaborate**: Join discussions and work on projects with a vibrant community of developers.\n\n`;
  content += `### Useful Links\n- [Trood Team Wiki](https://trood.com/teamspace): Access detailed documentation and resources to get started.\n`;
  content += `- [Join Trood Community](https://trood.com/launchpad): Connect with peers and explore opportunities to collaborate.\n\n`;

  // Contact section
  content += `### Contact Us\nIf you have any questions or want to learn more, feel free to reach out or join our community channels. Together, we can redefine what's possible in software development.\n\n`;

  return content;
}

// Main function
function updateReadme() {
  const submodules = parseGitmodules();
  const descriptions = loadDescriptions();
  const newReadmeContent = generateReadme(submodules, descriptions);

  // Write the new README content to README.md
  fs.writeFileSync(readmePath, newReadmeContent, "utf8");
}

// Run the update function
updateReadme();
