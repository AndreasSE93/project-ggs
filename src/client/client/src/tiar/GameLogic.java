package tiar;
/*
 * API:
 * 
 * 	#### Constructor
 * 	TiarGame()
 *   creates a new tic tac toe game (with an empty game field)
 *  
 *  
 *  
 *  #### Functions
 * 	makeMove(int position, int player)
 * 		puts player's sign in position if possible, and changes the turn. 1->2->1->2 etc.
 *  
 *  hasWon()
 *  	returns the player who has won, returns 0 if no1 has won.
 *  
 *  toString()
 *  	prints a representation of the gamefield
 *  
 *  
 *  isFull()
 *  	returns true if all the slots have been taken.
 *  
 *  validMove(int position, int player)
 *  	returns true if position is free and its player's turn.
 *  
 *  
 */








public class GameLogic{
	private int[] gameField = new int[9];
	private boolean isFull;
	private int turn;



	public GameLogic(){
		for (int i=0; i<9; i++){
			this.gameField[i] = 0;
		}

		this.isFull = false;
		this.turn = 1;
	}


	public int[] getGameField(){
		return this.gameField;
	}

	public void updateGameField(int position, int player){
		int[] field = getGameField();
		field[position] = player;	
		return;

	}


	public boolean IsFull(){
		return this.isFull;
	}

	public int getTurn(){
		return this.turn;
	}

	public void changeTurn(){
		if(this.getTurn() == 1){
			this.turn = 2;
		} else {this.turn = 1;}
	}


	public String toString(){
		int[] field = getGameField();
		System.out.println(field[0] + " " + field[1] + " " + field[2]);
		System.out.println(field[3] + " " + field[4] + " " + field[5]);
		System.out.println(field[6] + " " + field[7] + " " + field[8]);
		return("TEMP FIELD");
	}


	public void makeMove(int position, int player){

		if (validMove(position,player)){
			updateGameField(position,player);
			this.changeTurn();
		}
	}


	public boolean validMove(int position, int player){
		int[] field = getGameField();
		if(field[position] == 0 && this.getTurn() == player){return true;}
		else {return false;}
	}

	public int hasWon(){
		int[] field = getGameField();
		for (int i=1; i<3; i++){
			// Horizontal
			if (field[0] == field[1] && field[1] == field[2]){return field[0];}
			if (field[3] == field[4] && field[4] == field[5]){return field[3];}
			if (field[6] == field[7] && field[7] == field[8]){return field[6];}

			// Vertical
			if (field[0] == field[3] && field[3] == field[6]){return field[0];}
			if (field[1] == field[4] && field[4] == field[7]){return field[4];}
			if (field[2] == field[5] && field[5] == field[8]){return field[2];}

			// Diagonal
			if (field[0] == field[4] && field[4] == field[8]){return field[0];}
			if (field[2] == field[4] && field[4] == field[6]){return field[2];}

		}
		return 0;
	}

/*
	public static void main(String [] args){
		GameLogic Game = new GameLogic();
		Game.toString();
		System.out.println("PLAYING THE GAME FOR YOU");
		for (int i=0; i<9; i++){

			if (Game.getTurn() == 1){
				Game.makeMove(i,1);   	
			}
			else {Game.makeMove(i,2);}
		}

		Game.toString();
		if(Game.hasWon() == 1){System.out.println("Player 1 won");}
		if(Game.hasWon() == 2){System.out.println("Player 2 won");}
		if(Game.hasWon() == 0){System.out.println("no1 won :(");}





		GameLogic Game2 = new GameLogic();
		Game2.toString();
		System.out.println("PLAYING THE GAME FOR YOU");
		for (int i=0; i<9; i++){
			Game2.makeMove(i, 2);
		}

		Game2.toString();
		if(Game2.hasWon() == 1){System.out.println("Player 1 won");}
		if(Game2.hasWon() == 2){System.out.println("Player 2 won");}
		if(Game2.hasWon() == 0){System.out.println("no1 won :(");}





		GameLogic Game3 = new GameLogic();
		Game3.toString();
		System.out.println("PLAYING THE GAME FOR YOU");
		Game3.makeMove(0,1);	
		Game3.makeMove(1,2);
		Game3.makeMove(2,2);
		Game3.makeMove(3,2);
		Game3.makeMove(4,1);
		Game3.makeMove(5,1);
		Game3.makeMove(6,1);
		Game3.makeMove(7,1);
		Game3.makeMove(8,2);
		Game3.toString();
		if(Game3.hasWon() == 1){System.out.println("Player 1 won");}
		if(Game3.hasWon() == 2){System.out.println("Player 2 won");}
		if(Game3.hasWon() == 0){System.out.println("no1 won :(");}





	}

*/
}






