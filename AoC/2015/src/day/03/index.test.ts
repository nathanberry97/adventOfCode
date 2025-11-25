import { describe, it, expect } from 'vitest';
import { partOne, partTwo } from './index';

describe('Day 03', () => {
    describe('Part One', () => {
        it('Test 1 ~ to be 2', () => {
            // Arrange
            const input = '>';

            // Act
            const result = partOne(input);

            // Assert
            expect(result).toBe(2);
        });

        it('Test 2 ~ to be 4', () => {
            // Arrange
            const input = '^>v<';

            // Act
            const result = partOne(input);

            // Assert
            expect(result).toBe(4);
        });

        it('Test 3 ~ to be 2', () => {
            // Arrange
            const input = '^v^v^v^v^v';

            // Act
            const result = partOne(input);

            // Assert
            expect(result).toBe(2);
        });
    });

    describe('Part Two', () => {
        it('Test 1 ~ to be 3', () => {
            // Arrange
            const input = '^v';

            // Act
            const result = partTwo(input);

            // Assert
            expect(result).toBe(3);
        });

        it('Test 2 ~ to be 3', () => {
            // Arrange
            const input = '^>v<';

            // Act
            const result = partTwo(input);

            // Assert
            expect(result).toBe(3);
        });

        it('Test 3 ~ to be 11', () => {
            // Arrange
            const input = '^v^v^v^v^v';

            // Act
            const result = partTwo(input);

            // Assert
            expect(result).toBe(11);
        });
    });
});

