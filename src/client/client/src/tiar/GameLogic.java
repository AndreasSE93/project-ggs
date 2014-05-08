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


	public String toString(){
		int[] field = getGameField();
		System.out.println(field[0] + " " + field[1] + " " + field[2]);
		System.out.println(field[3] + " " + field[4] + " " + field[5]);
		System.out.println(field[6] + " " + field[7] + " " + field[8]);

		String board = field[0] + " " + field[1] + " " + field[2] +"/n" +
					 	field[3] + " " + field[4] + " " + field[5] +"/n" +
					 	field[6] + " " + field[7] + " " + field[8] +"/n";

		return(board);
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
			if (field[0] == field[1] && field[0] == field[2] && field[0] != 0){return field[0];}
			if (field[3] == field[4] && field[3] == field[5] && field[3] != 0){return field[3];}
			if (field[6] == field[7] && field[6] == field[8] && field[6] != 0){return field[6];}

			// Vertical
			if (field[0] == field[3] && field[0] == field[6] && field[0] != 0){return field[0];}
			if (field[1] == field[4] && field[1] == field[7] && field[1] != 0){return field[4];}
			if (field[2] == field[5] && field[2] == field[8] && field[2] != 0){return field[2];}

			// Diagonal
			if (field[0] == field[4] && field[0] == field[8] && field[0] != 0){return field[0];}
			if (field[2] == field[4] && field[2] == field[6] && field[2] != 0){return field[2];}
		}
		return 0;
	}

	public int isDraw(){
		int[] gamefield = this.getGameField();
		for (int i=0; i<9; i++){
			if (gamefield[i] == 0){return 0;}
		}
		return 1;
	}

}






