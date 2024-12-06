package puzzle

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testInput = `
		#####
		#.^.#
		##..#`

func TestParseMapIntoGrid(t *testing.T) {
	grid := parseMapIntoGrid(strings.TrimSpace(testInput))

	expectedRow0 := []string{"#", "#", "#", "#", "#"}
	expectedRow1 := []string{"#", ".", "^", ".", "#"}
	expectedRow2 := []string{"#", "#", ".", ".", "#"}

	assert.Equal(t, expectedRow0, grid[0])
	assert.Equal(t, expectedRow1, grid[1])
	assert.Equal(t, expectedRow2, grid[2])

	assert.Len(t, grid, 3)
}

func TestFindStartingPosition(t *testing.T) {
	grid := parseMapIntoGrid(strings.TrimSpace(testInput))
	startingPosition, err := findStartingPosition(grid)

	assert.NoError(t, err)
	assert.Equal(t, 1, startingPosition.Y)
	assert.Equal(t, 2, startingPosition.X)
}

func TestIsBlocked(t *testing.T) {
	grid := parseMapIntoGrid(strings.TrimSpace(testInput))
	startingPosition, _ := findStartingPosition(grid)

	actual := isBlocked(grid, startingPosition, &Point2D{X: 0, Y: 1})
	assert.False(t, actual)

	actual = isBlocked(grid, startingPosition, &Point2D{X: 0, Y: -1})
	assert.True(t, actual)

	actual = isBlocked(grid, startingPosition, &Point2D{X: 1, Y: 0})
	assert.False(t, actual)

	actual = isBlocked(grid, startingPosition, &Point2D{X: -1, Y: 0})
	assert.False(t, actual)

	actual = isBlocked(grid, startingPosition, &Point2D{X: -2, Y: 0})
	assert.True(t, actual)

	actual = isBlocked(grid, startingPosition, &Point2D{X: -3, Y: 0})
	assert.False(t, actual)
}

func TestMove(t *testing.T) {
	currentPosition := Point2D{X: 0, Y: 0}
	movingVector := Point2D{X: 0, Y: 1}

	newPosition := move(&currentPosition, &movingVector)
	assert.Equal(t, 0, newPosition.X)
	assert.Equal(t, 1, newPosition.Y)

	movingVector = Point2D{X: 0, Y: -1}
	newPosition = move(&currentPosition, &movingVector)
	assert.Equal(t, 0, newPosition.X)
	assert.Equal(t, -1, newPosition.Y)

	movingVector = Point2D{X: 1, Y: 0}
	newPosition = move(&currentPosition, &movingVector)
	assert.Equal(t, 1, newPosition.X)
	assert.Equal(t, 0, newPosition.Y)

	movingVector = Point2D{X: -1, Y: 0}
	newPosition = move(&currentPosition, &movingVector)
	assert.Equal(t, -1, newPosition.X)
	assert.Equal(t, 0, newPosition.Y)
}

func TestTurnRight(t *testing.T) {
	movingVector := Point2D{X: 0, Y: -1} // Move up

	movingVector = *turnRight(&movingVector) // Move right
	assert.Equal(t, 1, movingVector.X)
	assert.Equal(t, 0, movingVector.Y)

	movingVector = *turnRight(&movingVector) // Move down
	assert.Equal(t, 0, movingVector.X)
	assert.Equal(t, 1, movingVector.Y)

	movingVector = *turnRight(&movingVector) // Move left
	assert.Equal(t, -1, movingVector.X)
	assert.Equal(t, 0, movingVector.Y)

	movingVector = *turnRight(&movingVector) // Move up
	assert.Equal(t, 0, movingVector.X)
	assert.Equal(t, -1, movingVector.Y)
}

func TestIsOutOfBounds(t *testing.T) {
	grid := parseMapIntoGrid(strings.TrimSpace(testInput))

	currentLocation := Point2D{X: 0, Y: 0}
	actual := isOutOfBounds(grid, currentLocation)
	assert.False(t, actual)

	currentLocation = Point2D{X: 4, Y: 0}
	actual = isOutOfBounds(grid, currentLocation)
	assert.False(t, actual)

	currentLocation = Point2D{X: 0, Y: 2}
	actual = isOutOfBounds(grid, currentLocation)
	assert.False(t, actual)

	currentLocation = Point2D{X: 4, Y: 2}
	actual = isOutOfBounds(grid, currentLocation)
	assert.False(t, actual)

	currentLocation = Point2D{X: -1, Y: 0}
	actual = isOutOfBounds(grid, currentLocation)
	assert.True(t, actual)

	currentLocation = Point2D{X: 5, Y: 0}
	actual = isOutOfBounds(grid, currentLocation)
	assert.True(t, actual)

	currentLocation = Point2D{X: 0, Y: -1}
	actual = isOutOfBounds(grid, currentLocation)
	assert.True(t, actual)

	currentLocation = Point2D{X: 0, Y: 3}
	actual = isOutOfBounds(grid, currentLocation)
	assert.True(t, actual)
}

func TestCreatePositionKey(t *testing.T) {
	key := createPositionKey(1, 2)
	assert.Equal(t, "1,2", key)

	key = createPositionKey(2, 1)
	assert.Equal(t, "2,1", key)

	key = createPositionKey(-1, -2)
	assert.Equal(t, "-1,-2", key)
}

func TestCreatePositionWithFacingKey(t *testing.T) {
	key := createPositionWithFacingKey(1, 2, Point2D{X: 3, Y: 4})
	assert.Equal(t, "1,2,3,4", key)

	key = createPositionWithFacingKey(4, 3, Point2D{X: 2, Y: 1})
	assert.Equal(t, "4,3,2,1", key)

	key = createPositionWithFacingKey(-1, -2, Point2D{X: -3, Y: -4})
	assert.Equal(t, "-1,-2,-3,-4", key)
}

func TestMoveUntilOutOfBounds(t *testing.T) {
	grid := parseMapIntoGrid(strings.TrimSpace(testInput))
	actual := moveUntilOutOfBounds(grid)
	assert.Equal(t, 3, actual.UniqueLocationCount)
	assert.Equal(t, 3, actual.StepsUntilOutOfBounds)
	assert.ElementsMatch(t, actual.VisitedLocations, []Point2D{
		{X: 2, Y: 1},
		{X: 3, Y: 1},
		{X: 3, Y: 2},
	})

}
