package snake;

public class SnakePlayer {

	public String PlayerName;
	public int PlayerID;
	private double PosX;
	private double PosY;
	private boolean Alive;
	
	SnakePlayer(int player, double x, double y, boolean alive){
		PlayerID = player;
		PosX = x;
		PosY = y;
		Alive = alive;
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
	
	
}



