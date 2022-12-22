require_relative "../main.rb"

describe "GridPoint" do
  let(:x) { 4 }
  let(:y) { 7 }
  let(:obj) { GridPoint.new(x, y) }

  describe "#x" do
    it { expect(obj.x).to eq x }
  end

  describe "#y" do
    it { expect(obj.y).to eq y }
  end

  describe "#notation" do
    it "should return '(x,y)' as String" do
      expect(obj.notation).to eq "(#{x},#{y})"
    end
  end

  describe "#same_coordinates_with?" do
    context "when two GridPoint object have same coordinates" do
      it "should return true" do
        expect(obj.same_coordinates_with?(GridPoint.new(x, y))).to eq true
      end
    end

    context "when two GridPoint object have different coordinates" do
      it "should return true" do
        expect(obj.same_coordinates_with?(GridPoint.new(x+1, y+1))).to eq false
      end
    end
  end

  # 隣り合う条件 (x-1,y),(x+1,y),(x,y-1),(x,y+1)
  describe "#neighbor_of?" do
    it "(x-1,y) is neighbor" do
      expect(obj.neighbor_of?(GridPoint.new(x-1, y))).to eq true
    end
    it "(x+1,y) is neighbor" do
      expect(obj.neighbor_of?(GridPoint.new(x+1, y))).to eq true
    end
    it "(x,y-1) is neighbor" do
      expect(obj.neighbor_of?(GridPoint.new(x, y-1))).to eq true
    end
    it "(x,y+1) is neighbor" do
      expect(obj.neighbor_of?(GridPoint.new(x, y+1))).to eq true
    end
    it "(x+1,y+1) is not neighbor" do
      expect(obj.neighbor_of?(GridPoint.new(x+1, y+1))).to eq false
    end
  end
end

describe "GridPointSet" do
  let(:grid_point1) { GridPoint.new(1, 2) }
  let(:grid_point2) { GridPoint.new(3, 6) }
  let(:grid_point_set) { GridPointSet.new(grid_point1, grid_point2) }

  it do
    expect(grid_point_set.grid_points).to eq [grid_point1, grid_point2]
  end

  describe "#contain?" do
    context "when contains given grid_points contain" do
      it "should return true" do
        expect(grid_point_set.contain?(grid_point1)).to eq true
      end
    end
    context "when doesn't contains given grid_points contain" do
      it "should return false" do
        expect(grid_point_set.contain?(GridPoint.new(10, 20))).to eq false
      end
    end
  end

  # 隣り合う条件 (x-1,y),(x+1,y),(x,y-1),(x,y+1)
  describe "#connected?" do
    context "when (x-1,y) pattern" do
      let(:grid_point1) { GridPoint.new(1, 2) }
      let(:grid_point2) { GridPoint.new(0, 2) }
      it "should return true" do
        expect(grid_point_set.connected?).to eq true
      end
    end
    context "when (x+1,y) pattern" do
      let(:grid_point1) { GridPoint.new(1, 2) }
      let(:grid_point2) { GridPoint.new(2, 2) }
      it "should return true" do
        expect(grid_point_set.connected?).to eq true
      end
    end
    context "when (x,y-1) pattern" do
      let(:grid_point1) { GridPoint.new(1, 2) }
      let(:grid_point2) { GridPoint.new(1, 1) }
      it "should return true" do
        expect(grid_point_set.connected?).to eq true
      end
    end
    context "when (x,y+1) pattern" do
      let(:grid_point1) { GridPoint.new(1, 2) }
      let(:grid_point2) { GridPoint.new(1, 3) }
      it "should return true" do
        expect(grid_point_set.connected?).to eq true
      end
    end
  end
end
