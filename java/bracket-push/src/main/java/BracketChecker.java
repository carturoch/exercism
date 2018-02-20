import java.util.Stack;

public class BracketChecker {
    private String line;

    BracketChecker(String line) {
        this.line = line;
    }

    public boolean areBracketsMatchedAndNestedCorrectly() {
        Stack<Character> open = new Stack<>();
        for (char c :
                line.toCharArray()) {
            if (isOpener(c)) {
                open.push(c);
            }
            if (isCloser(c)) {
                if (open.empty()) {
                    return false;
                }
                if (doNotMatchOpen(c, open.pop())) {
                    return false;
                }
            }
        }
        return open.empty();
    }

    private boolean doNotMatchOpen(char c, char open) {
        if (c == ']' && open == '[') {
            return false;
        }
        if (c == ')' && open == '(') {
            return false;
        }
        if (c == '}' && open == '{') {
            return false;
        }

        return true;
    }

    private boolean isCloser(char c) {
        return c == '}' || c == ']' || c == ')';
    }

    private boolean isOpener(char c) {
        return c == '{' || c == '[' || c == '(';
    }

}
