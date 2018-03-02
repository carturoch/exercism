import java.util.List;
import java.util.function.Function;
import java.util.function.IntUnaryOperator;
import java.util.function.Supplier;
import java.util.stream.Collector;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

class LargestSeriesProductCalculator {

    private final List<Integer> digits;

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
        int upperBoundary = digits.size() - numberOfDigits + 1;
        final Function<Integer, IntUnaryOperator> subListProduct = steps -> start -> productOf(start, steps);

        return IntStream
                .range(0, upperBoundary)
                .map(subListProduct.apply(numberOfDigits))
                .reduce(0, Integer::max);
    }

    private int productOf(int start, int steps) {
        return digits
                .subList(start, start + steps)
                .stream()
                .reduce((n, acc) -> n * acc)
                .get();
    }
}
