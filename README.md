# Notes - Command-line tool
This project is a simple command-line tool to manage notes with features like adding, showing, and deleting notes. It uses SQLite as the database and is written in Go.

## Features

- **Add Notes**: Add notes to the database.
- **Show Notes**: View all saved notes.
- **Delete Notes**: Remove a specific note by its ID.

## Installation

1. Visit our [site]() to download the zip file
### Or
1.
```bash 
git clone https://github.com/gaurav1952/notes
cd notes
```
2. Install dependencies:
```bash
go mod tidy
```
3. Build the project:
(according to your machine OS)
```bash
#example for windows
go build -o notes.exe 
```

#### You can add this build file to your environment file 

To add the notes build file (i.e., notes.exe for Windows or notes for macOS/Linux) to your system's environment path, follow these instructions for each platform:

## 1. Windows
Adding notes.exe to **PATH**:
1. Move the notes.exe file to a permanent location **(e.g., C:\tools\notes)**:
   - Create a folder for your tools, **e.g., C:\tools\notes**
   - Move the **notes.exe** file there.
2. Add the folder to the system PATH:
   - Open the Start Menu and search for **Environment Variables**.
   - Select **Edit** the system environment variables.
   - In the System Properties window, click the **Environment Variables button**.
   - In the Environment Variables window, under System variables, select **Path** and click **Edit**.
   - Click New and add the path to the folder where notes.exe is located **(e.g., C:\tools\notes)**.
   - Click **OK** to close all windows.
3. Test the **PATH**:
   - **Open Command Prompt and run:**
  ```bash
notes.exe help
```
  - If the command runs successfully, the notes.exe file is now accessible from any directory.

### or 
**Run this command in PowerShell**
```powershell
$pathToAdd = "C:\path\to\your\notes.exe"
$env:Path += ";$pathToAdd"
[Environment]::SetEnvironmentVariable("Path", $env:Path, [EnvironmentVariableTarget]::User)
```

### Steps: 
1. Replace **C:\path\to\your\notes.exe** with the actual path where **notes.exe** is located **(e.g., C:\tools\notes)**.
2. **Open PowerShell as an administrator**.
3. Run the command above.
### Verify:
To verify that the path has been added, you can check the **Path** variable:
### Note: 
This command updates the **Path** for the current user. If you want to set it for the system (for all users), replace **[EnvironmentVariableTarget]::User** with **[EnvironmentVariableTarget]::Machine** in the last line, but note that you need administrative privileges to do this.


## 2. MacOS 
Adding **notes** to PATH:
1. Move the **notes** binary to a permanent location **(e.g., /usr/local/bin)**:

```bash
sudo mv notes /usr/local/bin/notes
```
2. Make the file executable:
```bash
sudo chmod +x /usr/local/bin/notes
```



3.**Test the PATH**: Open a terminal and run:
```bash
notes help
```
If the command runs successfully, the notes binary is now accessible from any directory.

### Note:
Ensure **/usr/local/bin** is included in your **PATH**. You can check this with:
```bash
 echo $PATH
```
## 3. Linux
Adding **notes** to PATH:
1. Move the notes binary to a location included in your **PATH** **(e.g., /usr/local/bin)**:

```bash
sudo mv notes /usr/local/bin/notes
```

2. Make the file executable:
```bash
sudo chmod +x /usr/local/bin/notes
```
3. **Test the PATH:** Open a terminal and run:
```bash 
notes help
```
If the command runs successfully, the notes binary is now accessible from any directory.

### Note:
Ensure **/usr/local/bin** is included in your **PATH**. You can check this with:
```bash
 echo $PATH
```


