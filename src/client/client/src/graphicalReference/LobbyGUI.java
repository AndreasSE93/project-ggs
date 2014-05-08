package graphicalReference;

import java.awt.*;
import java.io.File;
import java.io.IOException;
import java.util.ArrayList;

import javax.imageio.ImageIO;
import javax.swing.*;
import javax.swing.border.EmptyBorder;

import packageManaging.HostRoom;



public class LobbyGUI {
	public JPanel chatt = new JPanel();
	public JTextField field;
	public JTextPane chatPanel = new JTextPane();
	public JFrame lobby = new JFrame();
	public ChatGUI chatgui;
	public final String userName;
	public JButton joinButton;
	public JButton createButton;
	public JButton refreshButton;
	public JList<String> createList;
	public JList<HostRoom> joinList;
	
	public JPanel westTopPanel;
	public JPanel eastTopPanel;
	
	public LobbyGUI(String usr){
		this.userName = usr;
		
	}
	
	public void render(ArrayList<HostRoom> players, ArrayList<String> games) {

			
		
		lobby.setLayout(new BorderLayout());
		lobby.getContentPane().setBackground(Color.DARK_GRAY);
		lobby.setSize(800, 600);
		lobby.setLocationRelativeTo(null);
		lobby.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		lobby.setBackground(Color.DARK_GRAY);

		JPanel topPanel = new JPanel();
		JPanel bottomPanel = new JPanel();
		
		eastTopPanel = new JPanel();
		westTopPanel = new JPanel();
		
		JPanel centerTopPanel = new JPanel();
		topPanel.setLayout(new BorderLayout());
		eastTopPanel.setLayout(new BorderLayout());
		westTopPanel.setLayout(new BorderLayout());
		refreshButton = new JButton();
		makeJButton(refreshButton, "resources/refreshButton.png", "refreshbutton");
		centerTopPanel.setLayout(new BorderLayout());
		refreshButton.setBackground(Color.DARK_GRAY);
		centerTopPanel.add(refreshButton, BorderLayout.SOUTH);
		centerTopPanel.setBackground(Color.DARK_GRAY);
		topPanel.add(centerTopPanel, BorderLayout.CENTER);
		
		makePlayerList(eastTopPanel, players);
		makeGameList(westTopPanel, games);
		
		topPanel.add(westTopPanel, BorderLayout.WEST);
		topPanel.add(eastTopPanel, BorderLayout.EAST);
		topPanel.setBackground(Color.DARK_GRAY);
		Dimension preferredSize = new Dimension(400, 300);
		topPanel.setPreferredSize(preferredSize);
		lobby.add(topPanel, BorderLayout.NORTH);
		lobby.add(bottomPanel, BorderLayout.SOUTH);
		bottomPanel.add(chatt, "Center");
		chatgui = new ChatGUI(this.userName);
		chatgui.render(chatt, lobby);

		topPanel.setBorder(null);
		lobby.validate();
		lobby.setVisible(true);
		

	}

	public void makePlayerList(JPanel panel, ArrayList<HostRoom> L) {
		joinButton = new JButton();
		makeJButton(joinButton, "resources/JoinButton.png", "joinbutton");

		panel.add(joinButton, BorderLayout.SOUTH);
		addArrayList(panel,L);
		
	}
	
	public void makeGameList(JPanel panel, ArrayList<String> L) {
		createButton = new JButton();
		makeJButton(createButton,"resources/CreateButton.png", "createbutton");
		panel.add(createButton, BorderLayout.SOUTH);
		addArrayListString(panel,L);
		
	}
	
	public void addArrayList(JPanel panel, ArrayList<HostRoom> L){
		final JList<HostRoom> jList = getJList(L);

		jList.setBackground(Color.DARK_GRAY.darker());
		jList.setForeground(Color.LIGHT_GRAY);
		panel.setBorder(new EmptyBorder(10, 50, 10, 50));
		panel.setBackground(Color.DARK_GRAY);
		
		JScrollPane js = new JScrollPane(jList);
		js.setHorizontalScrollBar(null);
		js.setBorder(null);
		panel.add(js, BorderLayout.CENTER);
	
			joinList = jList;
		
			
		
		
		
	}
	public void addArrayListString(JPanel panel, ArrayList<String> L){
		final JList<String> jList = getJListString(L);

		jList.setBackground(Color.DARK_GRAY.darker());
		jList.setForeground(Color.LIGHT_GRAY);
		panel.setBorder(new EmptyBorder(10, 50, 10, 50));
		panel.setBackground(Color.DARK_GRAY);
		
		JScrollPane js = new JScrollPane(jList);
		js.setHorizontalScrollBar(null);
		js.setBorder(null);
		panel.add(js, BorderLayout.CENTER);
	
			createList = jList;
		
		
		
		
		
	}

	public static JList<String> getJListString(ArrayList<String> list) {
		int size = list.size();
		final DefaultListModel<String> model = new DefaultListModel<String>();
		for (int i = 0; i < size; i++) {
			model.addElement(list.get(i));
		}
		return new JList<String>(model);
	}


	public static JList<HostRoom> getJList(ArrayList<HostRoom> list) {
		int size = list.size();
		final DefaultListModel<HostRoom> model = new DefaultListModel<HostRoom>();
		for (int i = 0; i < size; i++) {
			model.addElement(list.get(i));
		}
		return new JList<HostRoom>(model);
	}
	
	public void makeJButton (JButton jb, String src, String actionCommand){
		File imageCheck = new File(src);
		try {
			Image img = ImageIO.read(imageCheck);
			jb.setIcon(new ImageIcon(img));
		} catch (IOException ex) {
			System.out.println("Hej");
		}
		jb.setMargin(new Insets(0, 0, 0, 0));
		jb.setBorder(null);
		jb.setActionCommand(actionCommand);
		
	}


}