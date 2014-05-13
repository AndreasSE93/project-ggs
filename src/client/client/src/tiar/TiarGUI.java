package tiar;

import graphicalReference.ChatGUI;

import java.awt.BorderLayout;
import java.awt.Color;
import java.awt.Dimension;
import java.awt.Font;
import java.awt.Graphics;
import java.awt.GridLayout;
import java.awt.Image;
import java.awt.Insets;
import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;

import javax.imageio.ImageIO;

import javax.swing.ImageIcon;
import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JTextArea;

import javax.swing.JPanel;
import javax.swing.border.EmptyBorder;

public class TiarGUI extends GameLogic {
	public JButton a1 = new JButton();
	public JButton a2 = new JButton();
	public JButton a3 = new JButton();
	public JButton b1 = new JButton();
	public JButton b2 = new JButton();
	public JButton b3 = new JButton();
	public JButton c1 = new JButton();
	public JButton c2 = new JButton();
	public JButton c3 = new JButton();
	public JButton[] game = new JButton[9];

	public JFrame window = new JFrame();
	public JPanel chatPanel = new JPanel();

	public ChatGUI chat;

	public GameLogic gl = new GameLogic();
	

	
	public void render(String usr) {

		window.setSize(1035,790);
		window.setLayout(new BorderLayout());
		
		window.setLocationRelativeTo(null);
		window.setResizable(false);
		window.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		window.setBackground(Color.DARK_GRAY);

		JPanel gamePane = new JPanel(new BorderLayout());
		gamePane.setPreferredSize(new Dimension(1035, 550));
		gamePane.setBackground(Color.BLACK);

		JPanel gameInfo = new JPanel();
		gameInfo.setPreferredSize(new Dimension(150,550));
		gameInfo.setBackground(Color.gray.darker());
		gamePane.add(gameInfo, BorderLayout.EAST);
		
		
		JTextArea gameName = new JTextArea("Tic Tac Toe");
		Font font = new Font("Helvetica", Font.BOLD, 12);
		gameName.setFont(font);
		gameName.setEditable(false);
		gameName.setBackground(Color.gray.darker());
		gameName.setForeground(Color.WHITE);
		gameInfo.add(gameName, BorderLayout.NORTH);
		
		
		JPanel gameCointainer = new JPanel();
		gameCointainer.setBackground(Color.DARK_GRAY.darker());
		gameCointainer.setBorder(new EmptyBorder(30,30,0,0));
		
		ImagePanel gameBoard = new ImagePanel();
		gameBoard.setLayout(new GridLayout(3,3));
		//gameBoard.setVisible(true);


		gameBoard.setBorder(new EmptyBorder(10,10,10,10));
		
		gameBoard.add(a1);
		gameBoard.add(a2);
		gameBoard.add(a3);
		gameBoard.add(b1);
		gameBoard.add(b2);
		gameBoard.add(b3);
		gameBoard.add(c1);
		gameBoard.add(c2);
		gameBoard.add(c3);
		
		a1.setName("a1");
		a2.setName("a2");
		a3.setName("a3");
		b1.setName("b1");
		b2.setName("b2");
		b3.setName("b3");
		c1.setName("c1");
		c2.setName("c2");
		c3.setName("c3");


		a1.setPreferredSize(new Dimension(150, 150));
		a2.setPreferredSize(new Dimension(150, 150));
		a3.setPreferredSize(new Dimension(150, 150));
		b1.setPreferredSize(new Dimension(150, 150));
		b2.setPreferredSize(new Dimension(150, 150));
		b3.setPreferredSize(new Dimension(150, 150));
		c1.setPreferredSize(new Dimension(150, 150));
		c2.setPreferredSize(new Dimension(150, 150));
		c3.setPreferredSize(new Dimension(150, 150));

		game[0] = a1;
		game[1] = a2;
		game[2] = a3;
		game[3] = b1;
		game[4] = b2;
		game[5] = b3;
		game[6] = c1;
		game[7] = c2;
		game[8] = c3;

		for(int i =0; i<game.length; i++){
			invisibleButton(game[i]);
			
		}

		gameCointainer.add(gameBoard);
		gamePane.add(gameCointainer);

		window.add(gamePane, BorderLayout.NORTH);
		chat = new ChatGUI(usr);
		chat.render(chatPanel, window);


		window.setVisible(true);
		window.validate();

	}

	public void updateGameBoard(int[] gameBoard){
		
			gl.gameField = gameBoard;
			File imageCross = new File("resources/cross_150_150.png");
		
			File imageCircle = new File("resources/circle_150_150.png");
			Image img1 = null;
			Image img2 = null;
			try {
				img1 = ImageIO.read(imageCross);
				img2 = ImageIO.read(imageCircle);
			} catch (IOException e) {
				// TODO Auto-generated catch block
				e.printStackTrace();
			}
		
		for (int i = 0; i<9; i++){
			if (gameBoard[i] == 1){
				
				this.game[i].setIcon(new ImageIcon(img1));
			}
			else if(gameBoard[i] == 2){
				this.game[i].setIcon(new ImageIcon(img2));
			}
			else this.game[i].setIcon(null);
		}
		
	}
	
	public void doMove(String position, int player) {

		int pos = getInt(position);
		if (gl.validMove(pos, player)) {
			JButton update = this.game[pos];
			gl.makeMove(pos, player);
			makeJButton(update, player);

		}

	}
	
	public void invisibleButton(JButton a){
		a.setOpaque(false);
		a.setContentAreaFilled(false);
		a.setBorderPainted(false);
		a.setBorder(new EmptyBorder(5,5,5,5));
	}
	
	public int getInt(String Pos) {
		switch (Pos) {
		case "a1":
			return 0;
		case "a2":
			return 1;
		case "a3":
			return 2;
		case "b1":
			return 3;
		case "b2":
			return 4;
		case "b3":
			return 5;
		case "c1":
			return 6;
		case "c2":
			return 7;
		case "c3":
			return 8;
		default:
			return 10000;
		}

	}

	private void makeJButton(JButton jb, int player) {
		File imageCheck;
		if (player == 1) {
			imageCheck = new File("resources/cross_150_150.png");
		} else {
			imageCheck = new File("resources/circle_150_150.png");
		}
		try {
			Image img = ImageIO.read(imageCheck);
			jb.setIcon(new ImageIcon(img));
		} catch (IOException ex) {
		}
		jb.setMargin(new Insets(0, 0, 0, 0));
		jb.setBorder(null);

	}

	public void clearBoard(){
		gl.clearBoard();
		for(int i=0; i< game.length; i++){
			game[i].setIcon(null); 
		}
		
	
	
	}
}

	class ImagePanel extends JPanel{

	    /**
		 * 
		 */
		private static final long serialVersionUID = 1L;
		private BufferedImage image;

	    public ImagePanel() {
	       try {                
	          image = ImageIO.read(new File("resources/tttBoard_150_150.png"));
	          } catch (IOException ex) {
	            // handle exception...
	       }
	    }

	    @Override
	    protected void paintComponent(Graphics g) {
	        super.paintComponent(g);
	        g.drawImage(image, 0, 0, null); // see javadoc for more info on the parameters            
	    }

}

