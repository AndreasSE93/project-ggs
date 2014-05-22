package snake;

import java.awt.BorderLayout;
import java.awt.Color;
import java.awt.Dimension;
import java.awt.Font;
import java.awt.Image;
import java.awt.image.BufferedImage;

import javax.swing.ImageIcon;
import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JLabel;
import javax.swing.JPanel;
import javax.swing.JTextArea;
import javax.swing.border.EmptyBorder;

public class SnakeGUI {

	public JFrame window = new JFrame();
	ImageIcon ii;
	BufferedImage bi;
	public JLabel gamePane;
	Image dimg;
	public JPanel achtungPanel;
	public JButton leaveGame;
	public JButton startGame;
	final private int WIDTH = 1035 - 200;
	final private int HEIGTH = 790;

	private JTextArea p1, p2, p3, p4;
	private JTextArea[] playerScores = { p1, p2, p3, p4 };

	private int[] ColorArray = { Color.red.getRGB(), Color.blue.getRGB(),
			Color.green.getRGB(), Color.yellow.getRGB() };

	public void render() {
		window.setSize(1035, 790);
		window.setLayout(new BorderLayout());

		window.setLocationRelativeTo(null);
		window.setResizable(false);
		window.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		window.setBackground(Color.DARK_GRAY);

		achtungPanel = new JPanel(new BorderLayout());
		achtungPanel.setPreferredSize(new Dimension(1035, 790));

		gamePane = new JLabel();
		gamePane.setFocusable(true);
		// gamePane.addKeyListener(this);

		achtungPanel.add(createTextPanel(), BorderLayout.EAST);
		achtungPanel.add(gamePane);

		bi = new BufferedImage(WIDTH, HEIGTH, BufferedImage.TYPE_INT_RGB);
		gamePane.setFocusable(true);
		renderNewGame();

		window.add(achtungPanel, BorderLayout.NORTH);
		
		window.setVisible(true);
		window.validate();

		// Temp play

		/*SnakePlayer[] players = new SnakePlayer[4];
		int posy1 = 100;
		int posy2 = 200;
		int posy3 = 300;
		int posy4 = 400;
		for (int i = 0; i < (1035 - 200); i++) {
			SnakePlayer p1 = new SnakePlayer(1, i, posy1, true, 0);
			SnakePlayer p2 = new SnakePlayer(2, i, posy2, true, 10);
			SnakePlayer p3 = new SnakePlayer(3, i, posy3, true, 15);
			SnakePlayer p4 = new SnakePlayer(4, i, posy4, true, i);
			players[0] = p1;
			players[1] = p2;
			players[2] = p3;
			players[3] = p4;
			repaint(players);
			try {
				Thread.sleep(15);
			} catch (InterruptedException e) {
				// TODO Auto-generated catch block
				e.printStackTrace();
			}

		}*/

	}

	public JPanel createTextPanel() {

		JPanel gameInfo = new JPanel();
		gameInfo.setPreferredSize(new Dimension(200, 790));
		gameInfo.setBackground(Color.gray.darker());

		JTextArea gameName = new JTextArea("Achtung Die Kurve");
		Font font = new Font("Helvetica", Font.BOLD, 12);
		gameName.setFont(font);
		gameName.setEditable(false);
		gameName.setBackground(Color.gray.darker());
		gameName.setForeground(Color.WHITE);
		gameName.setBorder(new EmptyBorder(0, 0, 200, 0));
		gameInfo.add(gameName, BorderLayout.NORTH);

		createPlayerPanels(gameInfo);
		JPanel buttons = new JPanel();

		leaveGame = new JButton("Leave");
		leaveGame.setFont(font);
		leaveGame.setBackground(Color.gray.darker());

		startGame = new JButton("Start");
		startGame.setFont(font);
		startGame.setBackground(Color.gray.darker());
		startGame.setActionCommand("startButton");

		buttons.setBorder(new EmptyBorder(250, 0, 0, 0));
		buttons.setBackground(Color.gray.darker());
		buttons.add(startGame);
		buttons.add(leaveGame);
		gameInfo.add(buttons, BorderLayout.SOUTH);

		return gameInfo;

	}

	public void createPlayerPanels(JPanel gameInfo) {
		Font playerFont = new Font("Helvetica", Font.BOLD, 20);
		int i;
		for (i = 0; i < playerScores.length; i++) {
			playerScores[i] = new JTextArea("Player " + i + ": 0");

			playerScores[i].setFont(playerFont);
			playerScores[i].setBackground(Color.gray.darker());
			playerScores[i].setBorder(new EmptyBorder(0, 0, 30, 0));
			Color clr = new Color(ColorArray[i]);
			playerScores[i].setForeground(clr);
			gameInfo.add(playerScores[i]);

		}
		// playerScores[i-1].setBorder(new EmptyBorder(0,0,250,0));
	}

	public void updateScore(SnakePlayer p) {
		int id = p.PlayerID - 1;
		String stringText = playerScores[id].getText();
		int index = 0;
		for (int i = stringText.length() - 1; i >= 0; i--) {
			if (stringText.substring(i-1,i).equals(":")) {
				index = i;
				break;
			}
		
		}
		
		playerScores[id].setText(stringText.substring(0, index) + " " + p.getScore());

	}

	public void repaint(SnakePlayer[] players) {
		for (int i = 0; i < players.length; i++) {
			// if(i == 0){ uncomment to only track player 1
			//System.out.println(i + "PosX:" + players[i].getPosX() + "PosY"
			//		+ players[i].getPosY());
			//System.out.println("Kom hit till paint");
			if (players[i].isAlive()) {

				updateScore(players[i]);
				for(int d =0; d<10; d=d+2){
				bi.setRGB((int) players[i].playerArray[d],
						(int) players[i].playerArray[d+1],
						ColorArray[players[i].PlayerID - 1]);
				ii = new ImageIcon();
				ii.setImage(bi);

				gamePane.setIcon(ii);
				gamePane.validate();
				}
				ii = new ImageIcon();
				ii.setImage(bi);

				gamePane.setIcon(ii);
				gamePane.validate();
				// }
			}
		}
	}

	public void renderNewGame() {

		/*
		 * for(int index= 0; index<PlayerArray.length; index++){
		 * PlayerArray[index].revivePlayer(); }
		 */

		for (int k = 0; k < WIDTH; k++) {
			for (int i = 0; i < HEIGTH; i++) {
				bi.setRGB(k, i, Color.BLACK.getRGB());

			}
		}
		ii = new ImageIcon();
		ii.setImage(bi);

		gamePane.setIcon(ii);
		gamePane.validate();

	}

	/*public static void main(String[] arg) {
		SnakeGUI SG = new SnakeGUI();
		SG.render();

	}*/

}
