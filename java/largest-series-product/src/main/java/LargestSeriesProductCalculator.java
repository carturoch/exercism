import java.util.List;
import java.util.stream.Collector;
import java.util.stream.Collectors;

class LargestSeriesProductCalculator {

    private List<Integer> digits;

    LargestSeriesProductCalculator(String inputNumber) throws IllegalArgumentException {
        if (inputNumber != "" && !inputNumber.matches("\\d+")) {
            throw new IllegalArgumentException("String to search may only contain digits.");
        }

        this.digits = inputNumber
                .chars()
                .mapToObj(Character::getNumericValue)
                .collect(Collectors.toList());
    }

    long calculateLargestProductForSeriesLength(int numberOfDigits) throws IllegalArgumentException {
        if (numberOfDigits < 0) {
            throw new IllegalArgumentException("Series length must be non-negative.");
        }
        if (digits.size() < numberOfDigits) {
            throw new IllegalArgumentException("Series length must be less than or equal to the length of the string to search.");
        }
        if (digits.size() == numberOfDigits && numberOfDigits == 0) {
            return 1;
        }
        long best = 0;
        int upperBoundary = digits.size() - numberOfDigits + 1;
        for (int i = 0; i < upperBoundary; i++) {
            long product = digits.get(i);
            for (int j = 1; j < numberOfDigits; j++) {
                product *= digits.get(i + j);
            }
            best = Long.max(best, product);
        }
        return best;
    }
}
