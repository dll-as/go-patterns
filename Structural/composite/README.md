# Composite Pattern (Go)

## What is Composite?

Composite is a structural design pattern.
It lets you **treat individual objects and compositions of objects uniformly**.

You use this pattern when:
- You have **tree-like structures**
- You want **uniform operations** on both single objects and groups
- You want to simplify client code

---

## Common Real-World Examples

- File system (File / Folder)
- UI components (Button, Panel)
- Organization charts (Employee, Department)
- Graphics (Shape, Group of Shapes)

---

## Example in this project

We have a **file system**:
- Files (leaf nodes)
- Folders (composite nodes)

Client can **print structure** without caring if it’s a file or folder.

---

## 🧠 Key Points (Simple)

* `File` is a leaf node
* `Folder` is a composite node containing children
* Client treats both uniformly via `FileSystem` interface
* Easy to add new types of nodes (e.g., Shortcut, Symlink)

---

## 🆚 Composite vs Decorator

| Composite                     | Decorator                             |
| ----------------------------- | ------------------------------------- |
| Tree structure                | Add behavior dynamically              |
| Leaf & Composite treated same | Wrap an object with new functionality |
| Example: File / Folder        | Example: Logging wrapper              |