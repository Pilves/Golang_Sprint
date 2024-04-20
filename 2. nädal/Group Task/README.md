# Note Organizer 📓

Welcome to the **Note Organizer** - a sleek and simple tool designed to help you manage your personal notes efficiently and effectively! Whether you're a student, a professional, or anyone in between, our tool is crafted to make note-taking as easy as pie 🥧.

## 🌟 What does the Note Organizer do?

**Note Organizer** is a command-line application that allows you to create, display, edit, and delete your notes in a text-based interface. Simply run the tool with the name of your note collection, and you're good to go! It's perfect for managing notes on the fly and ensures that your thoughts are organized and accessible anytime.

## 🔍 How to Use

Getting started with the **Note Organizer** is super easy! Here’s how you can get the most out of this tool:

### Starting the Tool

To start managing your notes, simply use the following command:

```bash
./notestool collection_name
```
Replace `collection_name` with your desired name for the note collection. This will be used as the filename where all of your notes are saved, making it easy to switch between different projects or contexts by simply changing the file name when you start the tool.

### Commands

Once started, the tool presents you with the following options:

- `1`: Display all notes
- `2`: Add a new note
- `3`: Remove an existing note
- `4`: Exit the application

### Examples

0. **Starting the program:**

```bash
$ ./notestool
Usage: ./notestool collection_name
collection_name: The name of the note collection to manage.
help: Display this help message.
```

1. **Adding a Note:**

```bash
$ ./notestool example
1. Display notes
2. Add a note
3. Remove a note
4. Exit
Enter your choice:
2
Enter the note text:
this is an example note

```
2. **Displaying Notes:**

```bash
Enter your choice:
1
001 - This is an example note
002 - This is another example note!
003 - This is third example note :)
```

3. **Removing a Note:**

```bash
Enter your choice:
3
Enter the number of the note to remove:
1

001 - This is another example note!
002 - This is third example note :)
```

## 💾 Data Storage

The **Note Organizer** uses a straightforward yet effective method to store your notes. All notes are stored in a plain text file, named according to the collection you specify. This file approach ensures that your notes are easily accessible and manageable:

- **File Format:** Each note is saved on a new line within the file. This format makes it straightforward to read and edit outside of the application if needed.
- **Real-Time Updates:** Whenever you add or remove a note, the file is updated instantly to reflect the changes, ensuring your data is always current without needing to perform manual saves.

## 🤓 Advanced Features

- **Persistent Storage:** Since your notes are stored in a text file, you can come back to them at any time, even after restarting your computer or moving to a different device, as long as you have access to the file.
- **Easy Navigation:** Simple command-line options for adding, deleting, and viewing notes help you manage your collections effortlessly.

## 🎨 Customizable Setup

You can personalize the Note Organizer to your liking by simply changing the text file's name you use for your notes collection. This flexibility allows you to have multiple collections for different projects or purposes.

Enjoy organizing your notes with the **Note Organizer**! 🚀📝
