package graphicalReference;

import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.io.File;
import java.io.IOException;
import java.util.ArrayList;


import javax.imageio.ImageIO;
import javax.swing.*;
import javax.swing.border.EmptyBorder;





public class LobbyGUI {
	public JPanel chatt = new JPanel();
	public JTextField field;
	public JTextPane chatPanel = new JTextPane();
    public JFrame lobby = new JFrame();
	public ChatGUI chatgui;
	public String userName;

public void render( ArrayList<String> players,  ArrayList<String> games) {
	
	
	

	lobby.setLayout(new BorderLayout());
	lobby.getContentPane().setBackground(Color.DARK_GRAY);
	lobby.setSize(800,600);
    lobby.setLocationRelativeTo(null);
    lobby.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
    lobby.setBackground(Color.DARK_GRAY);
    
    JPanel topPanel = new JPanel();
    JPanel bottomPanel = new JPanel();
    JPanel eastTopPanel = new JPanel();
    JPanel westTopPanel = new JPanel();
    topPanel.setLayout(new BorderLayout());
    eastTopPanel.setLayout(new BorderLayout());
    westTopPanel.setLayout(new BorderLayout());
    makeList(eastTopPanel, players,"gogogog");
    makeList(westTopPanel, games,"gogogog");
    topPanel.add(westTopPanel, BorderLayout.LINE_END);
    topPanel.add(eastTopPanel, BorderLayout.LINE_START);
    topPanel.setBackground(Color.DARK_GRAY);
    Dimension preferredSize = new Dimension(400,300);
    topPanel.setPreferredSize(preferredSize);
    lobby.add(topPanel, BorderLayout.NORTH);
    lobby.add(bottomPanel, BorderLayout.SOUTH);
    bottomPanel.add(chatt, "Center");
    chatgui = new ChatGUI();
    chatgui.render(chatt,lobby);
  
    
    
    topPanel.setBorder(null);
    lobby.validate();
    lobby.setVisible(true);
	
}



public void makeList(JPanel panel ,ArrayList<String> L, String nameButton){
	JButton game = new JButton();
	final JList<String> jList = getJList(L);
	//Dimension preferredSize = new Dimension(200 , 200);
	//jList.setPreferredSize(preferredSize);
	jList.setBackground(Color.DARK_GRAY.darker());
	jList.setForeground(Color.LIGHT_GRAY);
	panel.setBorder(new EmptyBorder(10,100,10,100));
	panel.setBackground(Color.DARK_GRAY);
	File imageCheck = new File("resources/connectButton.png");
	  try {
	    Image img = ImageIO.read(imageCheck);
	    game.setIcon(new ImageIcon(img));
	  } catch (IOException ex) {
	  }
	game.setMargin(new Insets(0, 0, 0, 0));
	game.setBorder(null);
	JScrollPane js = new JScrollPane(jList);
	js.setHorizontalScrollBar(null);
	js.setBorder(null);
	panel.add(game, BorderLayout.SOUTH);
	panel.add(js, BorderLayout.CENTER);
	game.addActionListener(new ActionListener(){
    	public void actionPerformed(ActionEvent ae){
    		String selectedItem = (String) jList.getSelectedValue();
    		System.out.println("Clicked" + selectedItem);
    	}
	});
}




public static JList<String> getJList(ArrayList<String> list){
	int size = list.size();
	final DefaultListModel<String> model = new DefaultListModel<String>();
	for (int i=0; i<size; i++){
		model.addElement(list.get(i));
	}
	return new JList<String>(model);
}
}