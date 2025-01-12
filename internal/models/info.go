package models

import "github.com/dasvh/go-learn-vim/internal/style"

// Section represents a section of information to be displayed on the Vim Information screen
type Section struct {
	Title   string
	Content string
}

const (
	tributeText = "Rest in peace, Bram Moolenaar, the creator of Vim, who passed away on August 3rd, 2023."
)

// BramTribute contains a tribute to Bram Moolenaar, the creator of Vim
var BramTribute = style.Styles.Display.SpecialText.Render(tributeText)

// VimInfoSections contains the sections of information to be displayed on the Vim Information screen
var VimInfoSections = []Section{
	{
		Title: "What is Vim?",
		Content: `Vim (Vi IMproved) is a powerful and efficient text editor, originally derived from the Vi editor in 1991.
It is widely used for its speed and ability to handle tasks with minimal effort,
making it a favorite tool for developers and system administrators.`,
	},
	{
		Title: "Why Use Vim?",
		Content: `Vim offers unmatched efficiency and productivity, making it a favorite among developers and system administrators.

Key features include:
- **Speed**: Optimized for keyboard-driven workflows, enabling faster text editing.
- **Modal Editing**: Switch between modes for navigation, editing, and selection, making it powerful and efficient.
- **Platform Independence**: Available on almost all systems, including POSIX and Unix-based platforms.
- **Customizability**: Personalize Vim with plugins, keybindings, and themes to suit your workflow.
- **Performance**: Designed to work in any environment, from lightweight terminals to advanced setups.
- **Availability**: Pre-installed on most Unix-based systems, ensuring it's accessible almost anywhere.`,
	},
	{
		Title: "Key Vim Motions",
		Content: `Vim motions are essential for navigating text efficiently. Here are some of the most important ones:
- **h**: Move left by one character.
- **l**: Move right by one character.
- **j**: Move down one line.
- **k**: Move up one line.
- **w**: Jump to the beginning of the next word.
- **b**: Jump to the beginning of the previous word.

Mastering these motions is the first step toward becoming proficient in Vim.`,
	},
	{
		Title: "The Philosophy of Vim:",
		Content: `Vim embraces the philosophy of efficiency through practice. It has a learning curve, but as you grow familiar 
with its commands and capabilities, it becomes a tool that adapts to your needs, making your workflow faster 
and more seamless.`,
	},
	{
		Title: "Additional information:",
		Content: `For more information on Vim, visit the official website at https://www.vim.org/ 
For a comprehensive guide on Vim, check out the Vim documentation at https://vimhelp.org/`,
	},
}

// CheatsheetSection contains the sections of information to be displayed on the Vim Cheatsheet screen
var CheatsheetSection = []Section{
	{
		Title: "Vim Modes",
		Content: `Vim operates in several modes, each with a specific purpose:
- **Normal Mode**: The default mode for navigation and running commands.
  - Press **ESC** to return to Normal Mode from any other mode.
- **Insert Mode**: For inserting and editing text.
  - Enter Insert Mode with **i** (insert before cursor), **a** (append after cursor), or **o** (open a new line).
- **Visual Mode**: For selecting and manipulating text.
  - Enter Visual Mode with **v** (character-wise selection), **V** (line-wise selection), or **Ctrl-v** (block selection).
- **Command Mode**: For executing advanced commands like saving, quitting, or searching.
  - Enter Command Mode with **:** (colon), then type commands like **:w** (save) or **:q** (quit).`,
	},
	{
		Title: "Basic Motions",
		Content: `- **h**: Move left by one character.
- **l**: Move right by one character.
- **j**: Move down one line.
- **k**: Move up one line.`,
	},
	{
		Title: "Advanced Motions",
		Content: `- **w**: Jump to the start of the next word.
- **e**: Jump to the end of the current/next word.
- **b**: Jump to the start of the previous word.`,
	},
	{
		Title: "Combining Motions with Numbers",
		Content: `- **[number][motion]**: Repeat a motion multiple times.
  - **2w**: Jump forward two words.
  - **3e**: Jump to the end of the next three words.
  - **4j**: Move down four lines.`,
	},
	{
		Title: "Using Motions with Commands",
		Content: `Motions can be combined with commands to perform actions:
- **d[motion]**: Delete up to the motion (e.g., **dw** deletes the current word).
- **y[motion]**: Copy (yank) up to the motion (e.g., **y2w** copies the next two words).
- **c[motion]**: Change (delete and switch to insert mode) up to the motion (e.g., **c3j** changes the next three lines).`,
	},
	{
		Title: "Essential Commands",
		Content: `- **:w**: Save the current file.
- **:q**: Quit Vim.
- **:wq**: Save and quit.
- **:q!**: Quit without saving changes.
- **ZZ**: Save and quit (shortcut).`,
	},
}
