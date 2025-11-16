import { describe, it, expect } from 'vitest';
import { partOne, partTwo } from './index';

describe('Day 01', () => {
    describe('Part One', () => {
        it('Test 1 ~ result to be 0', () => {
            // Arrange
            const input = '((()))';

            // Act
            const result = partOne(input)

            // Assert
            expect(result).toBe(0);
        });

        it('Test 2 ~ result to be 3', () => {
            // Arrange
            const input = '(()(()(';

            // Act
            const result = partOne(input)

            // Assert
            expect(result).toBe(3);
        });

        it('Test 3 ~ result to be -1', () => {
            // Arrange
            const input = '())';

            // Act
            const result = partOne(input)

            // Assert
            expect(result).toBe(-1);
        });

        it('Test 4 ~ result to be -3', () => {
            // Arrange
            const input = ')())())';

            // Act
            const result = partOne(input)

            // Assert
            expect(result).toBe(-3);
        });
    });

    describe('Part Two', () => {
        it('Test 1 ~ result to be 1', () => {
            // Arrange
            const input = ')';

            // Act
            const result = partTwo(input)

            // Assert
            expect(result).toBe(1);
        });

        it('Test 2 ~ result to be 5', () => {
            // Arrange
            const input = '()())';

            // Act
            const result = partTwo(input)

            // Assert
            expect(result).toBe(5);
        });
    });
});

