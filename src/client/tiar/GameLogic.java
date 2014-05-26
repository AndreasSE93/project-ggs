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
	public int[] gameField = new int[9];
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


	public void clearBoard(){
		int[] field = getGameField();
		for (int i=0; i<9; i++){field[i] = 0;}
		this.turn = 1;
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
	
}







