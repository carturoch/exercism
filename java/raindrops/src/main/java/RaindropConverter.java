import java.util.HashMap;
import java.util.Map;

class RaindropConverter {

    private static final Map<Integer, String> codes = new HashMap<Integer, String>(){{
        put(3, "Pling");
        put(5, "Plang");
        put(7, "Plong");
    }};

    String convert(int number) {
        String str = Integer.toString(number);
        Integer initialLength = str.length();
        for (Map.Entry<Integer, String> code :
                RaindropConverter.codes.entrySet()) {
            if (number % code.getKey() == 0) {
                str += code.getValue();
            }
        }
        if (initialLength == str.length()) {
            return str;
        }
        return str.substring(initialLength);
    }


}
