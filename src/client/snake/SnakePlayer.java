package snake;

public class SnakePlayer {

	public String PlayerName;
	public int PlayerID;
	private double PosX;
	private double PosY;
	private boolean Alive;
	public double []playerArray;
	private int Score;
	public SnakePlayer(int player, double x, double y, boolean alive, int s,double []pa){
		PlayerID = player;
		PosX = x;
		PosY = y;
		Alive = alive;
		Score = s;
		playerArray = pa;
	}
	
	public double getPosX(){
		return PosX;
	}
	
	public double getPosY(){
		return PosY;
	}

	
	public boolean isAlive(){
		return Alive;
	}
	
	public int getScore(){
		return Score;
	}
	
	
}



