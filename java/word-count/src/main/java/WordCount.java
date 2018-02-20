import java.util.HashMap;
import java.util.Map;

public class WordCount {
    public Map<String, Integer> phrase(String sentence) {
        String[] words = sentence
                .trim()
                .toLowerCase()
                .replaceAll("[^\\w']|\\B'\\b|\\b'\\B", " ")
                .split("\\s+");

        Map<String, Integer> totals = new HashMap<>();

        for (String w : words) {
            totals.put(w, totals.getOrDefault(w, 0) + 1);
        }
        return totals;
    }
}
