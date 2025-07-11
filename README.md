# JSON Formatter CLI

A lightweight command-line tool to validate, minify, and prettify JSON files.

---

### Install

1. **Download the binary for your system**  
   Download from the [Releases](https://github.com/nicolesoto-dev/json-formatter-cli) page (e.g., `json-formatter-cli_linux_amd64.tar.gz`) and extract it.

2. **Add alias to your shell config**  
   Open your shell configuration file (e.g., `.bashrc`, `.zshrc`) and add an alias:

   ```bash
   nano ~/.bashrc
   ```

3. **Add the following line (replace $PACKAGE_PATH with the actual binary path):**

   ```bash
   alias jsonfmt='$PACKAGE_PATH'
   ```

4. **Then reload your shell config:**

   ```bash
   source ~/.bashrc
   ```

5. **Use the CLI**
    ```bash
    jsonfmt --file $FILE_PATH
    ```

    Options:

    --output, -o: Output formatted JSON to a file

    --minify, -m: Minify the input JSON

    --prettify, -p: Prettify (indent) the input JSON

    --validate, -v: Validate JSON syntax only