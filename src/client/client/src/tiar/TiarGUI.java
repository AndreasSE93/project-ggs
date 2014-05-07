package tiar;

import java.awt.BorderLayout;
import java.awt.Color;
import java.awt.Dimension;
import java.util.ArrayList;


import javax.swing.JFrame;
import javax.swing.JPanel;

public class TiarGUI {
	public JPanel a1 = new JPanel();
	public JPanel a2 = new JPanel();
	public JPanel a3 = new JPanel();
	public JPanel b1 = new JPanel();
	public JPanel b2 = new JPanel();
	public JPanel b3 = new JPanel();
	public JPanel c1 = new JPanel();
	public JPanel c2 = new JPanel();
	public JPanel c3 = new JPanel();
	
	
	public  void render (ArrayList<String> gameField){
		
		JFrame window = new JFrame();
		window.setSize(800, 600);
		window.setLayout(new BorderLayout());
		window.setLocationRelativeTo(null);
		window.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		window.setBackground(Color.DARK_GRAY);
		
		JPanel gamePane = new JPanel();
		gamePane.setPreferredSize(new Dimension(500, 600));
		gamePane.setBackground(Color.BLACK);
		
		JPanel topRow = new JPanel();
		JPanel middleRow = new JPanel();
		JPanel bottomRow = new JPanel();
		
		topRow.setLayout(new BorderLayout());
		middleRow.setLayout(new BorderLayout());
		bottomRow.setLayout(new BorderLayout());
		
		topRow.add(this.a1, BorderLayout.LINE_START);
		topRow.add(this.a2, BorderLayout.AFTER_LAST_LINE);
		topRow.add(this.a3, BorderLayout.AFTER_LAST_LINE);
		
		middleRow.add(this.b1, BorderLayout.LINE_START);
		middleRow.add(this.b2, BorderLayout.AFTER_LAST_LINE);
		middleRow.add(this.b3, BorderLayout.AFTER_LAST_LINE);
		
		bottomRow.add(this.c1, BorderLayout.LINE_START);
		bottomRow.add(this.c2, BorderLayout.AFTER_LAST_LINE);
		bottomRow.add(this.c3, BorderLayout.AFTER_LAST_LINE);
		
		window.add(topRow);
		window.add(middleRow);
		window.add(bottomRow);
		
		
		
		
		
		window.add(gamePane, BorderLayout.WEST);
		window.setVisible(true);
		
		
		
		
	}
	
	
	public static void main (String args[]){
		TiarGUI t = new TiarGUI( );
	t.render(new ArrayList<String>());
	}
}
