import { describe, it, expect } from 'vitest';
import { partOne, partTwo } from './index';

describe('Day 02', () => {
    describe('Part One', () => {
        it('Test 1 ~ to be 58', () => {
            // Arrange
            const input = '2x3x4';

            // Act
            const result = partOne(input);

            // Assert
            expect(result).toBe(58);
        });

        it('Test 2 ~ to be 42', () => {
            // Arrange
            const input = '1x1x10';

            // Act
            const result = partOne(input);

            // Assert
            expect(result).toBe(43);
        });
    });

    describe('Part Two', () => {
        it('Test 1 ~ to be 34', () => {
            // Arrange
            const input = '2x3x4';

            // Act
            const result = partTwo(input);

            // Assert
            expect(result).toBe(34);
        });

        it('Test 2 ~ to be 14', () => {
            // Arrange
            const input = '1x1x10';

            // Act
            const result = partTwo(input);

            // Assert
            expect(result).toBe(14);
        });
    })
});

