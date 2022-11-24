class ParkingSystem {
    private int[] capacity = new int[3];
    private int[] parking = { 0, 0, 0 };

    public ParkingSystem(int big, int medium, int small) {
        capacity[0] = big;
        capacity[1] = medium;
        capacity[2] = small;
    }

    public boolean addCar(int carType) {
        if (parking[carType - 1] + 1 <= capacity[carType - 1]) {
            parking[carType - 1] += 1;
            return true;
        }

        return false;
    }
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * ParkingSystem obj = new ParkingSystem(big, medium, small);
 * boolean param_1 = obj.addCar(carType);
 */