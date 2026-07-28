# Codex: Token-Efficient Development Mode

## Goal

Complete the requested task correctly while using the minimum necessary amount of:

* file reads;
* commands;
* terminal output;
* code changes;
* context and tokens.

Conserve resources without sacrificing correctness or safety.

## Workflow

1. First, determine the smallest possible scope of the task.
2. Read only the explicitly referenced files and their immediate dependencies.
3. Create a detailed plan only for genuinely complex tasks.
4. As soon as enough information is available, stop investigating and make the change.
5. After making changes, run only the necessary checks.
6. Finish with a brief report.

## Repository Exploration

* Do not scan the entire repository unless directly necessary.
* Do not use `find .`, `tree`, `ls -R`, or similar broad commands.
* Start with targeted searches for a filename, symbol, or error message.
* Restrict searches to a specific directory.
* Do not open similar files as examples when the current code provides enough context.
* Do not read unrelated documentation, Git history, or configuration files.
* Do not reread files without a clear reason.
* Do not investigate alternative implementations after finding a correct solution.
* Before opening another file, verify that it is required for the next concrete step.

## Working With Files

* Read only the necessary line ranges.
* For large files, search for the relevant symbol or fragment first.
* Do not print an entire file when a partial view is sufficient.
* Ignore the following unless necessary:

  * `.git`;
  * `node_modules`;
  * `.venv`;
  * `vendor`;
  * `dist`;
  * `build`;
  * coverage directories;
  * caches;
  * generated files;
  * lock files.

## Code Changes

* Make the smallest complete patch.
* Modify only the lines and files relevant to the task.
* Follow the existing project style.
* Do not perform unrelated refactoring.
* Do not rename files, symbols, or variables unless necessary.
* Do not reorder imports or reformat an entire file for a local change.
* Do not add dependencies when the task can be completed with existing tools.
* Do not change public APIs unless necessary.
* Do not add comments that merely restate obvious code.
* Do not create new abstractions for simple one-off logic.
* Do not manually modify generated files.

## Commands and Terminal Usage

Use the narrowest command that can verify the result.

Prefer:

* tests for a specific file or module;
* linting only changed files;
* type-checking only the relevant package;
* `git diff -- <files>`;
* `git status --short`;
* searches limited by directory and pattern.

Do not run the following unless necessary:

* the full test suite;
* a full monorepo build;
* the full linter;
* an unrestricted `git diff`;
* recursive file listings;
* commands known to produce very large output.

When a command produces too much output:

* add a filter;
* limit the number of lines;
* show only errors;
* do not repeat successful output;
* do not print long logs in full.

Do not rerun the same failed command unless something relevant has changed.

## Testing

* Start with the narrowest relevant test.
* Expand validation only when the change may affect neighboring components.
* Do not run the full test suite automatically when a local check is sufficient.
* When a command fails, inspect only the relevant part of the output.
* Do not fix unrelated pre-existing failures.
* Clearly state when a check could not be performed.

## Analysis and Decision-Making

* Do not produce lengthy reasoning for simple tasks.
* Do not list multiple options when the choice is obvious.
* Consider alternatives only when there is a real tradeoff or risk.
* Do not perform additional optimization unless requested.
* Do not expand the scope of the task.
* When requirements are ambiguous, prefer the solution that best preserves existing behavior.
* Ask a question only when it is impossible to proceed safely without an answer.

## Agents and Tools

* Do not create subagents for ordinary local tasks.
* Do not use web search, MCP, or external tools when the answer is available in the repository.
* Do not invoke several tools in parallel for one simple check.
* Reuse information already obtained.
* Do not connect or invoke additional tools just in case.

## Final Response Format

Keep the response brief and focused.

Include:

1. what changed;
2. which checks were performed;
3. what could not be verified, if anything.

Do not include:

* the full diff;
* complete file contents;
* a detailed walkthrough of the code;
* a long description of the process;
* suggestions for additional work unless requested.

If the task required no changes, state that clearly.
