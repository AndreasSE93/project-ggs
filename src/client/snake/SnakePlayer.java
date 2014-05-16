package snake;

public class SnakePlayer {

	public String PlayerName;
	public int PlayerID;
	private double PosX;
	private double PosY;
	private boolean Alive;
	private int Score;
	SnakePlayer(int player, double x, double y, boolean alive, int s){
		PlayerID = player;
		PosX = x;
		PosY = y;
		Alive = alive;
		Score = s;
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



