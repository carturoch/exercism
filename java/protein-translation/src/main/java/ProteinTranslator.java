import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

class ProteinTranslator {

    private final static Map<String, String> codons = new HashMap<String, String>() {
        {
            put("AUG", "Methionine");
            put("UUU", "Phenylalanine");
            put("UUC", "Phenylalanine");
            put("UUA", "Leucine");
            put("UUG", "Leucine");
            put("UCU", "Serine");
            put("UCC", "Serine");
            put("UCA", "Serine");
            put("UCG", "Serine");
            put("UAU", "Tyrosine");
            put("UAC", "Tyrosine");
            put("UGU", "Cysteine");
            put("UGC", "Cysteine");
            put("UGG", "Tryptophan");
        }
    };

    private final static String[] stopCodons = {"UAA", "UAG", "UGA"};

    List<String> translate(String rnaSequence) {
        String seq = getUsableSequence(rnaSequence);
        return IntStream
                .range(0, seq.length())
                .filter(n -> n % 3 == 0)
                .mapToObj(i -> seq.substring(i, i + 3))
                .map(codons::get)
                .collect(Collectors.toList());
    }

    private String getUsableSequence(String rnaSequence) {
        String[] seqs = rnaSequence.split(String.join("|", stopCodons));
        return seqs.length == 0 ? "" : seqs[0];
    }
}